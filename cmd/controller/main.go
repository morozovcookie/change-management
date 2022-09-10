package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

func main() {
	cfg := NewConfig()
	if err := cfg.Parse(); err != nil {
		log.Fatalln(err)
	}

	logger, err := createLogger(cfg.logLevel)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(logger *zap.Logger) {
		if syncErr := logger.Sync(); syncErr != nil {
			log.Fatalln(syncErr)
		}
	}(logger)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	be := newBackend(cfg, logger)

	logger.Info("starting application")

	if err := be.init(ctx); err != nil {
		logger.Fatal("init", zap.Error(err))
	}

	eg.Go(func() error {
		logger.Info("starting http server", zap.String("host", be.apiServer.Addr))

		if err := be.apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("listen and serve http server: %w", err)
		}

		return nil
	})

	logger.Info("starting change request scheduler")
	crqDone := be.changeRequestScheduler.Schedule(ctx)

	logger.Info("application started")

	<-ctx.Done()

	logger.Info("stopping application")

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	if err := be.apiServer.Shutdown(ctx); err != nil {
		logger.Error("stopping http server", zap.Error(err))
	}

	be.pgxClient.Close()

	<-crqDone

	if err := eg.Wait(); err != nil {
		logger.Error("waiting for application be stopped", zap.Error(err))
	}

	logger.Info("application stopped")
}

func createLogger(logLevel zapcore.Level) (*zap.Logger, error) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(logLevel)

	logger, err := loggerConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("create logger: %w", err)
	}

	return logger.Named("controller"), nil
}
