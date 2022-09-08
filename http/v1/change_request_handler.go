package v1

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	cm "github.com/morozovcookie/change-management"
	"io"
	"net/http"
)

const (
	ChangeRequestHandlerPathPrefix = "/api/v1/change-requests"
	CreateChangeRequestPathPrefix  = "/"
)

var _ http.Handler = (*ChangeRequestHandler)(nil)

type ChangeRequestHandler struct {
	http.Handler

	changeRequestSvc cm.ChangeRequestService
}

func NewChangeRequestHandler(changeRequestSvc cm.ChangeRequestService) *ChangeRequestHandler {
	var (
		router  = chi.NewRouter()
		handler = &ChangeRequestHandler{
			Handler: router,

			changeRequestSvc: changeRequestSvc,
		}
	)

	router.Post(CreateChangeRequestPathPrefix, handler.handleCreateChangeRequest)

	return handler
}

type CreateChangeRequestRequest struct {
	RawType     string `json:"type"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	IsAutoClose bool   `json:"isAutoClose"`

	requestType cm.ChangeRequestType
}

func decodeCreateChangeRequestRequest(_ context.Context, reader io.Reader) (*CreateChangeRequestRequest, error) {
	decoded := new(CreateChangeRequestRequest)

	if err := json.NewDecoder(reader).Decode(decoded); err != nil {
		return nil, &cm.Error{
			Code:    cm.ErrorCodeInvalid,
			Message: "failed to decode CreateChangeRequest request",
			Err:     err,
		}
	}

	if decoded.requestType = cm.ChangeRequestType(decoded.RawType); !decoded.requestType.IsValid() {
		return nil, &cm.Error{
			Code:    cm.ErrorCodeInvalid,
			Message: `"type" parameter has an invalid value: ` + decoded.RawType,
			Err:     nil,
		}
	}

	return decoded, nil
}

type CreateChangeRequestResponse struct {
	ID string `json:"id"`
}

func newCreateChangeRequestResponse(crq *cm.ChangeRequest) *CreateChangeRequestResponse {
	return &CreateChangeRequestResponse{
		ID: crq.ID.String(),
	}
}

func (h *ChangeRequestHandler) handleCreateChangeRequest(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	decoded, err := decodeCreateChangeRequestRequest(ctx, request.Body)
	if err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	crq := &cm.ChangeRequest{
		Type:        decoded.requestType,
		Summary:     decoded.Summary,
		Description: decoded.Description,
		IsAutoClose: decoded.IsAutoClose,
	}

	if err := h.changeRequestSvc.CreateChangeRequest(ctx, crq); err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	encodeResponse(ctx, writer, http.StatusOK, newCreateChangeRequestResponse(crq))
}
