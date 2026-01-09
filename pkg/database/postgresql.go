package database

import (
	"go-api-starter/pkg/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do/v2"
)

type Postgresql struct {
	pool *pgxpool.Pool
}

func NewPostgresql(injector do.Injector) (*Postgresql, error) {
	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Postgresql

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s pool_max_conns=%d pool_min_conns=%d pool_max_conn_lifetime=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
		cfg.MaxOpenConns,
		cfg.MaxIdleConns,
		time.Duration(cfg.ConnMaxLifetime)*time.Second,
	)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	poolConfig.MaxConns = int32(cfg.MaxOpenConns)
	poolConfig.MinConns = int32(cfg.MaxIdleConns)
	poolConfig.MaxConnLifetime = time.Duration(cfg.ConnMaxLifetime) * time.Second
	poolConfig.HealthCheckPeriod = 1 * time.Minute
	poolConfig.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgresql{pool: pool}, nil
}

func (db *Postgresql) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *Postgresql) HealthCheckWithContext(ctx context.Context) error {
	if err := db.pool.Ping(ctx); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}
	return nil
}

func (db *Postgresql) Shutdown(context.Context) error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}
