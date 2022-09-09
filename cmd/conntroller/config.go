package main

import (
	"fmt"
	"net/url"
	"os"
)

type PgxConfig struct {
	Dsn *url.URL

	prefix string
}

func newPgxConfig(prefix string) *PgxConfig {
	return &PgxConfig{
		Dsn: &url.URL{
			Scheme: "postgres",
			User:   url.UserPassword("controller", "controller"),
			Host:   "localhost:5432",
			Path: "/cm?search_path=controller&pool_max_conns=10&pool_min_conns=5&pool_max_conn_lifetime=1h" +
				"&pool_max_conn_idle_time=30m&pool_health_check_period=1m",
		},

		prefix: prefix,
	}
}

func (cfg *PgxConfig) parse() error {
	if err := cfg.parseDsn(); err != nil {
		return fmt.Errorf("parse pgx config: %w", err)
	}

	return nil
}

func (cfg *PgxConfig) parseDsn() error {
	raw := os.Getenv(cfg.prefix + "DSN")
	if raw == "" {
		return nil
	}

	dsn, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("parse pgx dsn: %w", err)
	}

	cfg.Dsn = dsn

	return nil
}

type Config struct {
	Pgx *PgxConfig

	prefix string
}

func NewConfig() *Config {
	const prefix = "CONTROLLER_"
	return &Config{
		prefix: prefix,

		Pgx: newPgxConfig(prefix + "PGX_"),
	}
}

func (cfg *Config) Parse() error {
	for _, fn := range []func() error{
		cfg.Pgx.parse,
	} {
		if err := fn(); err != nil {
			return fmt.Errorf("parse config: %w", err)
		}
	}

	return nil
}
