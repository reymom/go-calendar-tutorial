package dao

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

type Config struct {
	ConnectionStringRead  string
	ConnectionStringWrite string
	WriteEnabled          bool
	MaxReadConnections    uint
	MaxWriteConnections   uint
}

type PsqlDao struct {
	readPool     *pgxpool.Pool
	writePool    *pgxpool.Pool
	readEnabled  bool
	writeEnabled bool
}

var _ model.TasksDao = (*PsqlDao)(nil)

func NewPsqlDao(config *Config) (*PsqlDao, error) {
	if config == nil {
		return nil, model.ErrNilNotAllowed
	}
	readPool, e := initializePool(config.ConnectionStringRead, config.MaxReadConnections)
	if e != nil {
		return nil, e
	}

	var writePool *pgxpool.Pool
	switch config.WriteEnabled {
	case true:
		writePool, e = initializePool(config.ConnectionStringWrite, config.MaxWriteConnections)
		if e != nil {
			return nil, e
		}
	case false:
		writePool, e = initializeEmptyPool(model.ErrFakeWritePool)
		if e != nil {
			return nil, e
		}
	}

	return &PsqlDao{
		readPool:     readPool,
		writePool:    writePool,
		readEnabled:  true,
		writeEnabled: config.WriteEnabled,
	}, nil
}

func initializePool(connectionString string, maxPoolSize uint) (*pgxpool.Pool, error) {

	poolConfig, e := pgxpool.ParseConfig(connectionString)
	if e != nil {
		return nil, e
	}
	poolConfig.MaxConns = int32(maxPoolSize)

	//We error after 11 seconds if no connection could be established
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	pool, e := pgxpool.NewWithConfig(ctx, poolConfig)
	if e != nil {
		return nil, e
	}
	e = pool.Ping(context.Background())
	if e != nil {
		return nil, e
	}

	return pool, nil
}

func initializeEmptyPool(fakePoolError error) (*pgxpool.Pool, error) {
	const fakeConnectionString = "postgresql://no-user:password@nonhost:1/non-existing-db"
	poolConfig, e := pgxpool.ParseConfig(fakeConnectionString)
	if e != nil {
		return nil, e
	}
	poolConfig.BeforeConnect = func(ctx context.Context, config *pgx.ConnConfig) error {
		return fakePoolError
	}

	pool, e := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if e != nil {
		return nil, e
	}

	return pool, nil
}
