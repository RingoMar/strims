package main

import (
	"errors"
	"fmt"
	"path"

	"github.com/MemeLabs/go-ppspp/pkg/kv"
	"github.com/MemeLabs/go-ppspp/pkg/kv/bbolt"
	"github.com/MemeLabs/go-ppspp/pkg/kv/postgres"
	"github.com/MemeLabs/go-ppspp/pkg/pathutil"
)

func openDB(cfg *Config) (kv.BlobStore, error) {
	switch cfg.Storage.Adapter.Get("bbolt") {
	case "bbolt":
		return bboltStorageAdapter(cfg)
	case "postgres":
		return postgresStorageAdapter(cfg)
	default:
		return nil, fmt.Errorf("unsupported storage adapter: %s", cfg.Storage.Adapter)
	}
}

func bboltStorageAdapter(cfg *Config) (kv.BlobStore, error) {
	dbPath, err := pathutil.Resolve(cfg.Storage.BBolt.Path.Get(path.Join("~", ".strims")))
	if err != nil {
		return nil, err
	}
	return bbolt.NewStore(dbPath)
}

func postgresStorageAdapter(cfg *Config) (kv.BlobStore, error) {
	connStr := cfg.Storage.Postgres.ConnStr.Get("")
	if connStr == "" {
		return nil, errors.New("postgres conn string empty")
	}
	return postgres.NewStore(connStr)
}