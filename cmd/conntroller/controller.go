package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/morozovcookie/change-management/http/v1"
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

	router := chi.NewRouter()
	router.Use(middleware.RealIP, middleware.RequestID, middleware.Recoverer, middleware.Logger)
	router.Mount(v1.ChangeRequestHandlerPathPrefix, v1.NewChangeRequestHandler())
	router.Mount(v1.IncidentHandlerPathPrefix, v1.NewIncidentHandler())

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       time.Second * 30,
		ReadHeaderTimeout: time.Second * 30,
		WriteTimeout:      time.Second * 30,
		IdleTimeout:       time.Minute * 3,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	logger.Info("starting application")

	eg.Go(func() error {
		logger.Info("starting http server", zap.String("host", srv.Addr))

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("stopping http server", zap.Error(err))
	}

	if err := eg.Wait(); err != nil {
		logger.Error("waiting for application be stopped", zap.Error(err))
	}

	logger.Info("application stopped")
}
