package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/samber/do/v2"
	"go-api-starter/core/config"
	"go-api-starter/core/logger"
	"time"
)

const (
	PostgresqlDriver = "postgres"
)

type Postgresql struct {
	db   *sql.DB
	sqlx *sqlx.DB
}

type Config struct {
	Host                   string
	Port                   int
	User                   string
	Password               string
	Database               string
	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetime        int    // in minutes
	SSLMode                string // disable, require, verify-ca, verify-full
	ConnectTimeout         int    // in seconds
	StatementTimeout       int    // in seconds
	IdleInTxSessionTimeout int    // in seconds
}

func NewPostgresql(injector do.Injector) (*Postgresql, error) {
	// Get configuration from the injector
	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Postgresql

	dsn := fmt.Sprintf(
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

	db, err := sqlx.Connect(PostgresqlDriver, dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Postgresql{db: db.DB}, nil
}

type IPostgresql interface {
	ExecContext(ctx context.Context, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	SQLx() *sqlx.DB
}

func (p *Postgresql) ExecContext(ctx context.Context, query string, args ...any) error {
	_, err := p.sqlx.ExecContext(ctx, query, args...)
	return err
}

func (p *Postgresql) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	return p.sqlx.GetContext(ctx, dest, query, args...)
}

func (p *Postgresql) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	return p.sqlx.SelectContext(ctx, dest, query, args...)
}

func (p *Postgresql) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

func (p *Postgresql) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}

func (p *Postgresql) NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error) {
	return p.sqlx.NamedQueryContext(ctx, query, arg)
}

func (p *Postgresql) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	return p.sqlx.NamedExecContext(ctx, query, arg)
}

func (p *Postgresql) SQLx() *sqlx.DB {
	return p.sqlx
}
