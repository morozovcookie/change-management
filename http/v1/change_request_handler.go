package v1

import (
	"context"
	"encoding/json"
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

func encodeResponse(_ context.Context, writer http.ResponseWriter, status int, response interface{}) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")

	if response == nil {
		return
	}

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}