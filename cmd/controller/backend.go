package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	cm "github.com/morozovcookie/change-management"
	v1 "github.com/morozovcookie/change-management/http/v1"
	"github.com/morozovcookie/change-management/nanoid"
	"github.com/morozovcookie/change-management/pgx"
	"go.uber.org/zap"
)

type backend struct {
	cfg    *Config
	logger *zap.Logger

	pgxClient *pgx.Client

	identifierGenerator cm.IdentifierGenerator

	changeRequestService cm.ChangeRequestService

	apiRouter chi.Router
	apiServer *http.Server
}

func newBackend(cfg *Config, logger *zap.Logger) *backend {
	apiRouter := chi.NewRouter()

	return &backend{
		cfg:    cfg,
		logger: logger,

		pgxClient: pgx.NewClient(cfg.Pgx.Dsn.String(), logger.Named("pgx")),

		identifierGenerator: nanoid.NewIdentifierGenerator(),

		apiRouter: apiRouter,
		apiServer: &http.Server{
			Addr:              cfg.HTTP.Address,
			Handler:           apiRouter,
			ReadTimeout:       cfg.HTTP.ReadTimeout,
			ReadHeaderTimeout: cfg.HTTP.ReadHeaderTimeout,
			WriteTimeout:      cfg.HTTP.WriteTimeout,
			IdleTimeout:       cfg.HTTP.IdleTimeout,
			MaxHeaderBytes:    cfg.HTTP.MaxHeaderBytes,
		},
	}
}

func (be *backend) init(ctx context.Context) error {
	if err := be.pgxClient.Connect(ctx); err != nil {
		return fmt.Errorf("init backend: %w", err)
	}

	be.initServices(ctx)
	be.setupAPIRoutes(ctx)

	return nil
}

func (be *backend) setupAPIRoutes(_ context.Context) {
	be.apiRouter.Use(middleware.RealIP, middleware.RequestID, middleware.Recoverer, middleware.Logger)
	be.apiRouter.Mount(v1.ChangeRequestHandlerPathPrefix, v1.NewChangeRequestHandler(be.changeRequestService))
	be.apiRouter.Mount(v1.IncidentHandlerPathPrefix, v1.NewIncidentHandler())
}

func (be *backend) initServices(ctx context.Context) {
	ii := []func(context.Context){
		be.initChangeRequestService,
	}

	quit := make(chan struct{}, 1)
	defer close(quit)

	for _, fn := range ii {
		go func(ctx context.Context, quit chan<- struct{}, fn func(context.Context)) {
			fn(ctx)
			quit <- struct{}{}
		}(ctx, quit, fn)
	}

	for range ii {
		<-quit
	}
}

func (be *backend) initChangeRequestService(_ context.Context) {
	be.changeRequestService = pgx.NewChangeRequestService(be.pgxClient, be.identifierGenerator)
}
