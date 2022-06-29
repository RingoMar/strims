// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/MemeLabs/strims/internal/dao"
	"github.com/MemeLabs/strims/internal/frontend"
	"github.com/MemeLabs/strims/internal/network"
	"github.com/MemeLabs/strims/internal/session"
	"github.com/MemeLabs/strims/pkg/apis/type/key"
	"github.com/MemeLabs/strims/pkg/httputil"
	"github.com/MemeLabs/strims/pkg/vnic"
	"github.com/MemeLabs/strims/pkg/vpn"
	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func runCmd(fs Flags) error {
	cfg, err := loadConfig(fs.String("config"))
	if err != nil {
		return err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	dao.Logger = logger

	var eg errgroup.Group
	var closers []io.Closer

	if cfg.Debug.Address.Ok() {
		eg.Go(func() error {
			srv := &http.Server{
				Addr:    cfg.Debug.Address.MustGet(),
				Handler: http.DefaultServeMux,
			}
			closers = append(closers, srv)
			logger.Debug("debug server starting", zap.String("address", cfg.Debug.Address.MustGet()))
			err := srv.ListenAndServe()
			logger.Debug("debug server exited", zap.Error(err))
			return err
		})
	}

	if cfg.Metrics.Address.Ok() {
		eg.Go(func() error {
			mux := http.NewServeMux()
			mux.Handle("/metrics", promhttp.Handler())
			srv := &http.Server{
				Addr:    cfg.Metrics.Address.MustGet(),
				Handler: mux,
			}
			closers = append(closers, srv)
			logger.Debug("metrics server starting", zap.String("address", cfg.Metrics.Address.MustGet()))
			err := srv.ListenAndServe()
			logger.Debug("metrics server exited", zap.Error(err))
			return err
		})
	}

	webRTCOpts := &vnic.WebRTCDialerOptions{
		ICEServers: cfg.VNIC.WebRTC.ICEServers.Get(nil),
		PortMin:    cfg.VNIC.WebRTC.PortMin,
		PortMax:    cfg.VNIC.WebRTC.PortMax,
		HostIP:     fs.String("host-ip"),
	}
	if cfg.VNIC.WebRTC.Enabled.Get(true) {
		if cfg.VNIC.WebRTC.UDPMuxAddress.Ok() {
			mux, lis, err := vnic.NewWebRTCUDPMux(cfg.VNIC.WebRTC.UDPMuxAddress.MustGet())
			if err != nil {
				return fmt.Errorf("creating webrtc udp mux: %w", err)
			}
			logger.Debug("webrtc udp mux started", zap.Stringer("address", lis.LocalAddr()))
			closers = append(closers, lis)
			webRTCOpts.UDPMux = mux
		}
		if cfg.VNIC.WebRTC.TCPMuxAddress.Ok() {
			mux, lis, err := vnic.NewWebRTCTCPMux(cfg.VNIC.WebRTC.TCPMuxAddress.MustGet(), cfg.VNIC.WebRTC.TCPReadBufferSize.Get(8))
			if err != nil {
				return fmt.Errorf("creating webrtc tcp mux: %w", err)
			}
			logger.Debug("webrtc tcp mux started", zap.Stringer("address", lis.Addr()))
			closers = append(closers, lis)
			webRTCOpts.TCPMux = mux
		}
	}

	httpmux := httputil.NewMapServeMux()
	newVPN := func(key *key.Key) (*vpn.Host, error) {
		var opts []vnic.HostOption
		if cfg.VNIC.Label.Ok() {
			opts = append(opts, vnic.WithLabel(cfg.VNIC.Label.MustGet()))
		}
		if cfg.VNIC.WebRTC.Enabled.Get(true) {
			opts = append(opts, vnic.WithInterface(vnic.NewWebRTCInterface(vnic.NewWebRTCDialer(logger, webRTCOpts))))
		}
		if cfg.VNIC.WebSocket.Enabled.Get(true) {
			opts = append(opts, vnic.WithInterface(vnic.NewWSInterface(logger, vnic.WSInterfaceOptions{
				ServeMux: httpmux,
			})))
		}
		host, err := vnic.New(logger, key, opts...)
		if err != nil {
			return nil, err
		}
		return vpn.New(logger, host)
	}

	store, err := openDB(cfg)
	if err != nil {
		return err
	}
	closers = append(closers, store)

	queue, err := openQueue(cfg)
	if err != nil {
		return err
	}
	closers = append(closers, queue)

	sessionManager := session.NewManager(logger, store, queue, newVPN, network.NewBroker(logger), httpmux)

	for _, s := range cfg.Session.Headless {
		_, err := sessionManager.GetOrCreateSession(s.ID, s.Key)
		if err != nil {
			return err
		}
	}

	if cfg.Session.Remote.Enabled.Get(false) {
		srv := &frontend.Server{
			Store:          store,
			Logger:         logger,
			SessionManager: sessionManager,
		}

		httpmux.HandleFunc("/api", session.KeyHandler(func(ctx context.Context, c *websocket.Conn) {
			err := srv.Listen(ctx, httputil.NewWSReadWriter(c))
			logger.Debug("remote client closed", zap.Error(err))
		}))
	}

	if cfg.HTTP.Address.Ok() {
		eg.Go(func() (err error) {
			srv := &http.Server{
				Addr:    cfg.HTTP.Address.MustGet(),
				Handler: httpmux,
			}
			closers = append(closers, srv)

			logger.Debug("app server starting", zap.String("address", cfg.HTTP.Address.MustGet()))
			if cfg.HTTP.TLS.Cert.Ok() && cfg.HTTP.TLS.Key.Ok() {
				err = srv.ListenAndServeTLS(cfg.HTTP.TLS.Cert.MustGet(), cfg.HTTP.TLS.Key.MustGet())
			} else {
				err = srv.ListenAndServe()
			}
			logger.Debug("app server exited", zap.Error(err))
			return err
		})
	}

	eg.Go(func() error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		err := fmt.Errorf("signal: %s", <-sig)

		for _, c := range closers {
			c.Close()
		}
		return err
	})

	return eg.Wait()
}

func addProfileCmd(fs Flags) error {
	cfg, err := loadConfig(fs.String("config"))
	if err != nil {
		return err
	}

	store, err := openDB(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	id, key, err := dao.CreateServerAuthThing(store, fs.String("username"), fs.String("password"))
	if err != nil {
		return err
	}
	log.Println(id)
	log.Println(base64.StdEncoding.EncodeToString(key))

	return nil
}
