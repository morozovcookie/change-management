package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	IncidentHandlerPathPrefix = "/api/v1/incidents"
	CreateIncidentPathPrefix  = "/"
)

var _ http.Handler = (*IncidentHandler)(nil)

type IncidentHandler struct {
	http.Handler
}

func NewIncidentHandler() *IncidentHandler {
	var (
		router  = chi.NewRouter()
		handler = &IncidentHandler{
			Handler: router,
		}
	)

	router.Post(CreateIncidentPathPrefix, handler.handleCreateIncident)

	return handler
}

func (h *IncidentHandler) handleCreateIncident(writer http.ResponseWriter, request *http.Request) {
	encodeResponse(request.Context(), writer, http.StatusOK, nil)
}
