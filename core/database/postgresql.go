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

type Database struct {
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

func NewDatabase(injector do.Injector) (*Database, error) {
	// Get configuration from the injector
	appConfig := do.MustInvoke[*config.Config](injector)
	cfg := appConfig.Database

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

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{db: db.DB}, nil
}

type IDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	SQLx() *sqlx.DB
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...any) error {
	_, err := d.sqlx.ExecContext(ctx, query, args...)
	return err
}

func (d *Database) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	return d.sqlx.GetContext(ctx, dest, query, args...)
}

func (d *Database) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	return d.sqlx.SelectContext(ctx, dest, query, args...)
}

func (d *Database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *Database) NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error) {
	return d.sqlx.NamedQueryContext(ctx, query, arg)
}

func (d *Database) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	return d.sqlx.NamedExecContext(ctx, query, arg)
}

func (d *Database) SQLx() *sqlx.DB {
	return d.sqlx
}
