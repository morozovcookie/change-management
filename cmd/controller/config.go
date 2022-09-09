package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
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
			Path: "/cm?pool_max_conns=10&pool_min_conns=5&pool_max_conn_lifetime=1h&pool_max_conn_idle_time=30m" +
				"&pool_health_check_period=1m",
		},

		prefix: prefix,
	}
}

func (cfg *PgxConfig) parse() error {
	if err := parseURL(cfg.prefix+"DSN", cfg.Dsn); err != nil {
		return fmt.Errorf("parse pgx config: %w", err)
	}

	return nil
}

func parseURL(name string, out any) error {
	if out == nil {
		panic(`"out" parameter could not be nil`)
	}

	ptr, ok := out.(*url.URL)
	if !ok {
		panic(`"out" parameter must be a *url.URL`)
	}

	val := os.Getenv(name)
	if val == "" {
		return nil
	}

	u, err := url.Parse(val)
	if err != nil {
		return fmt.Errorf(`parse "%s" variable: %w`, name, err)
	}

	*ptr = *u

	return nil
}

type HTTPConfig struct {
	Address           string
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int

	prefix string
}

func newHTTPConfig(prefix string) *HTTPConfig {
	return &HTTPConfig{
		Address:           "127.0.0.1:8080",
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
		IdleTimeout:       time.Minute * 3,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,

		prefix: prefix,
	}
}

func (cfg *HTTPConfig) parse() error {
	for _, pp := range []struct {
		name    string
		out     any
		parseFn func(string, any) error
	}{
		{
			name:    cfg.prefix + "ADDRESS",
			out:     &cfg.Address,
			parseFn: parseString,
		},
		{
			name:    cfg.prefix + "READ_TIMEOUT",
			out:     &cfg.ReadTimeout,
			parseFn: parseDuration,
		},
		{
			name:    cfg.prefix + "READ_HEADER_TIMEOUT",
			out:     &cfg.ReadHeaderTimeout,
			parseFn: parseDuration,
		},
		{
			name:    cfg.prefix + "WRITE_TIMEOUT",
			out:     &cfg.WriteTimeout,
			parseFn: parseDuration,
		},
		{
			name:    cfg.prefix + "IDLE_TIMEOUT",
			out:     &cfg.IdleTimeout,
			parseFn: parseDuration,
		},
		{
			name:    cfg.prefix + "MAX_HEADER_BYTES",
			out:     &cfg.MaxHeaderBytes,
			parseFn: parseInt,
		},
	} {
		if err := pp.parseFn(pp.name, pp.out); err != nil {
			return fmt.Errorf("parse HTTP config: %w", err)
		}
	}

	return nil
}

func parseString(name string, out any) error {
	if out == nil {
		panic(`"out" parameter could not be nil`)
	}

	ptr, ok := out.(*string)
	if !ok {
		panic(`"out" parameter must be a *string`)
	}

	val := os.Getenv(name)
	if val == "" {
		return nil
	}

	*ptr = val

	return nil
}

func parseDuration(name string, out any) error {
	if out == nil {
		panic(`"out" parameter could not be nil`)
	}

	ptr, ok := out.(*time.Duration)
	if !ok {
		panic(`"out" parameter must be a *time.Duration`)
	}

	val := os.Getenv(name)
	if val == "" {
		return nil
	}

	d, err := time.ParseDuration(val)
	if err != nil {
		return fmt.Errorf(`parse "%s" variable: %w`, name, err)
	}

	*ptr = d

	return nil
}

func parseInt(name string, out any) error {
	if out == nil {
		panic(`"out" parameter could not be nil`)
	}

	ptr, ok := out.(*int)
	if !ok {
		panic(`"out" parameter must be a *int`)
	}

	val := os.Getenv(name)
	if val == "" {
		return nil
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return fmt.Errorf(`parse "%s" variable: %w`, name, err)
	}

	*ptr = num

	return nil
}

type Config struct {
	Pgx      *PgxConfig
	HTTP     *HTTPConfig
	Monitor  *HTTPConfig
	Profile  *HTTPConfig
	logLevel zapcore.Level

	prefix string
}

func NewConfig() *Config {
	const prefix = "CONTROLLER_"
	return &Config{
		Pgx:      newPgxConfig(prefix + "PGX_"),
		HTTP:     newHTTPConfig(prefix + "HTTP_"),
		Monitor:  newHTTPConfig(prefix + "MONITOR_"),
		Profile:  newHTTPConfig(prefix + "PROFILE_"),
		logLevel: zapcore.ErrorLevel,

		prefix: prefix,
	}
}

func (cfg *Config) Parse() error {
	for _, fn := range []func() error{
		cfg.Pgx.parse,
		cfg.HTTP.parse,
		cfg.Monitor.parse,
		cfg.Profile.parse,
		cfg.parse,
	} {
		if err := fn(); err != nil {
			return fmt.Errorf("parse config: %w", err)
		}
	}

	return nil
}

func (cfg *Config) parse() error {
	if err := parseLogLevel(cfg.prefix+"LOG_LEVEL", &cfg.logLevel); err != nil {
		return fmt.Errorf("parse application config: %w", err)
	}

	return nil
}

func parseLogLevel(name string, out *zapcore.Level) error {
	if out == nil {
		panic(`"out" parameter could not be nil`)
	}

	val := os.Getenv(name)
	if val == "" {
		return nil
	}

	lvl, err := zapcore.ParseLevel(val)
	if err != nil {
		return fmt.Errorf(`parse "%s" variable: %w`, name, err)
	}

	*out = lvl

	return nil
}
