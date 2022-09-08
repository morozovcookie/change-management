package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	logger, err := loggerConfig.Build()
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

	be := newBackend()
	be.init(ctx)

	logger.Info("starting application")

	eg.Go(func() error {
		logger.Info("starting http server", zap.String("host", be.apiServer.Addr))

		if err := be.apiServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	logger.Info("application started")

	select {
	case <-ctx.Done():
		break
	}

	logger.Info("stopping application")

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	if err := be.apiServer.Shutdown(ctx); err != nil {
		logger.Error("stopping http server", zap.Error(err))
	}

	if err := eg.Wait(); err != nil {
		logger.Error("waiting for application be stopped", zap.Error(err))
	}

	logger.Info("application stopped")
}
