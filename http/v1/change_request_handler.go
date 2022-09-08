package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	ChangeRequestHandlerPathPrefix = "/api/v1/change-requests"
	CreateChangeRequestPathPrefix  = "/"
)

var _ http.Handler = (*ChangeRequestHandler)(nil)

type ChangeRequestHandler struct {
	http.Handler
}

func NewChangeRequestHandler() *ChangeRequestHandler {
	var (
		router  = chi.NewRouter()
		handler = &ChangeRequestHandler{
			Handler: router,
		}
	)

	router.Post(CreateChangeRequestPathPrefix, nil)

	return handler
}

func (h *ChangeRequestHandler) handleCreateChangeRequest(writer http.ResponseWriter, request *http.Request) {
	encodeResponse(request.Context(), writer, http.StatusOK, nil)
}
