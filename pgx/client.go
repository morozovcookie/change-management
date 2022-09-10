package pgx

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Queryer reads data from database.
type Queryer interface {
	// Query executes sql with args. It is safe to attempt to read from the
	// returned Rows even if an error is returned. The error will be the
	// available in rows.Err() after rows are closed. So it is allowed to
	// ignore the error returned from Query and handle it in Rows.
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)

	// QueryRow is a convenience wrapper over Query. Any error that occurs
	// while querying is deferred until calling Scan on the returned Row.
	// That Row will error with ErrNoRows if no rows are returned.
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

// Execer executes sql query.
type Execer interface {
	// Exec executes sql. sql can be either a prepared statement name or an SQL
	// string. arguments should be referenced positionally from the sql string
	// as $1, $2, etc.
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// QueryExecer reads data from database and executes sql query.
type QueryExecer interface {
	Queryer
	Execer
}

// TxBeginner starts a new transaction.
type TxBeginner interface {
	// BeginTx starts a transaction with txOptions determining the transaction
	// mode. Unlike database/sql, the context only affects the begin command.
	// i.e. there is no auto-rollback on context cancellation.
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

var (
	_ QueryExecer = (*Client)(nil)
	_ TxBeginner  = (*Client)(nil)
)

type Client struct {
	dsn    string
	logger *zap.Logger

	pool *pgxpool.Pool
}

func NewClient(dsn string, logger *zap.Logger) *Client {
	return &Client{
		dsn:    dsn,
		logger: logger,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer cancel()

	cfg, err := pgxpool.ParseConfig(c.dsn)
	if err != nil {
		return fmt.Errorf("connect to postgres: %w", err)
	}

	cfg.ConnConfig.Logger = zapadapter.NewLogger(c.logger)

	if c.pool, err = pgxpool.ConnectConfig(ctx, cfg); err != nil {
		return fmt.Errorf("connect to postgres: %w", err)
	}

	return nil
}

func (c *Client) Close() {
	c.pool.Close()
}

// Query executes sql with args. It is safe to attempt to read from the
// returned Rows even if an error is returned. The error will be the
// available in rows.Err() after rows are closed. So it is allowed to
// ignore the error returned from Query and handle it in Rows.
func (c *Client) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.pool.Query(ctx, sql, args...) //nolint:wrapcheck
}

// QueryRow is a convenience wrapper over Query. Any error that occurs
// while querying is deferred until calling Scan on the returned Row.
// That Row will error with ErrNoRows if no rows are returned.
func (c *Client) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return c.pool.QueryRow(ctx, sql, args...)
}

// Exec executes sql. sql can be either a prepared statement name or an SQL
// string. arguments should be referenced positionally from the sql string
// as $1, $2, etc.
func (c *Client) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return c.pool.Exec(ctx, sql, arguments...) //nolint:wrapcheck
}

// BeginTx starts a transaction with txOptions determining the transaction
// mode. Unlike database/sql, the context only affects the begin command.
// i.e. there is no auto-rollback on context cancellation.
func (c *Client) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return c.pool.BeginTx(ctx, txOptions) //nolint:wrapcheck
}
