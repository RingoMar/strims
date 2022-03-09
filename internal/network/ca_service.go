package network

import (
	"context"
	"errors"
	"sync"

	"github.com/MemeLabs/go-ppspp/internal/dao"
	networkv1 "github.com/MemeLabs/go-ppspp/pkg/apis/network/v1"
	networkv1ca "github.com/MemeLabs/go-ppspp/pkg/apis/network/v1/ca"
	"github.com/MemeLabs/go-ppspp/pkg/apis/type/certificate"
	"go.uber.org/zap"
)

// AddressSalt ...
var AddressSalt = []byte("ca")

// New ...
func newCAService(logger *zap.Logger, store *dao.ProfileStore, network *networkv1.Network) *CAService {
	return &CAService{
		logger:   logger,
		store:    store,
		logCache: dao.NewCertificateLogCache(store, nil),
		network:  network,
	}
}

// CAService ...
type CAService struct {
	logger   *zap.Logger
	store    *dao.ProfileStore
	logCache dao.CertificateLogCache
	lock     sync.Mutex
	network  *networkv1.Network

	// invite policy
	// certificate revocation stream
	// certificate transparency list?
}

// Close ...
func (s *CAService) Close() {
	s.logCache.Close()
}

// Renew ...
func (s *CAService) Renew(ctx context.Context, req *networkv1ca.CARenewRequest) (*networkv1ca.CARenewResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	err := dao.VerifyCertificateRequest(req.CertificateRequest, certificate.KeyUsage_KEY_USAGE_PEER|certificate.KeyUsage_KEY_USAGE_SIGN)
	if err != nil {
		return nil, err
	}

	// TODO: verify invitation policy
	// TODO: verify nick
	cert, err := dao.SignCertificateRequestWithNetwork(req.CertificateRequest, s.network.GetServerConfig())
	if err != nil {
		return nil, err
	}

	log, err := dao.NewCertificateLog(s.store, s.network.Id, cert)
	if err != nil {
		return nil, err
	}
	err = dao.CertificateLogs.Insert(s.store, log)
	if err != nil {
		return nil, err
	}
	s.logCache.Store(log)

	return &networkv1ca.CARenewResponse{Certificate: cert}, nil
}

// Find ...
func (s *CAService) Find(ctx context.Context, req *networkv1ca.CAFindRequest) (*networkv1ca.CAFindResponse, error) {
	if req.Subject == "" && req.SerialNumber == nil {
		return nil, errors.New("find request must specify subject or serial number")
	}

	var log *networkv1ca.CertificateLog
	var err error
	if req.Subject != "" {
		log, err = s.logCache.BySubject.Get(dao.FormatCertificateLogSubjectKey(s.network.Id, req.Subject))
	} else {
		log, err = s.logCache.BySerialNumber.Get(dao.FormatCertificateLogsSerialNumberKey(s.network.Id, req.SerialNumber))
	}
	if err != nil {
		return nil, err
	}

	if req.FullChain {
		cert := log.Certificate
		for cert.GetParentSerialNumber() != nil {
			parentLog, err := s.logCache.BySerialNumber.Get(dao.FormatCertificateLogsSerialNumberKey(s.network.Id, cert.GetParentSerialNumber()))
			if err != nil {
				return nil, err
			}
			cert.ParentOneof = &certificate.Certificate_Parent{Parent: parentLog.Certificate}
			cert = parentLog.Certificate
		}
	}

	return &networkv1ca.CAFindResponse{Certificate: log.Certificate}, nil
}
