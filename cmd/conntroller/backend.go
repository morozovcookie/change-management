package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	cm "github.com/morozovcookie/change-management"
	v1 "github.com/morozovcookie/change-management/http/v1"
	"github.com/morozovcookie/change-management/pgx"
)

type backend struct {
	changeRequestService cm.ChangeRequestService

	apiRouter chi.Router
	apiServer *http.Server
}

func newBackend() *backend {
	apiRouter := chi.NewRouter()

	return &backend{
		apiRouter: apiRouter,
		apiServer: &http.Server{
			Addr:              ":8080",
			Handler:           apiRouter,
			ReadTimeout:       time.Second * 30,
			ReadHeaderTimeout: time.Second * 30,
			WriteTimeout:      time.Second * 30,
			IdleTimeout:       time.Minute * 3,
			MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		},
	}
}

func (be *backend) init(ctx context.Context) {
	be.initServices(ctx)
	be.setupAPIRoutes(ctx)
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
	be.changeRequestService = pgx.NewChangeRequestService()
}
