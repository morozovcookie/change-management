package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	cm "github.com/morozovcookie/change-management"
)

const (
	ChangeRequestHandlerPathPrefix = "/api/v1/change-requests"
	CreateChangeRequestPathPrefix  = "/"
	GetChangeRequestPathPrefix     = "/{changeRequestId}"
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
	router.Get(GetChangeRequestPathPrefix, handler.handleGetChangeRequest)

	return handler
}

type CreateChangeRequestRequest struct {
	IsAutoClose bool   `json:"isAutoClose"`
	RawType     string `json:"type"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	ExternalID  string `json:"externalId"`

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

	crq, err := h.changeRequestSvc.FindChangeRequestByExternalID(ctx, decoded.ExternalID)
	if err != nil && cm.ErrorCodeFromError(err) != cm.ErrorCodeNotFound {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	if crq != nil {
		encodeResponse(ctx, writer, http.StatusOK, newCreateChangeRequestResponse(crq))

		return
	}

	crq = &cm.ChangeRequest{
		Type:        decoded.requestType,
		Summary:     decoded.Summary,
		Description: decoded.Description,
		IsAutoClose: decoded.IsAutoClose,
		ExternalID:  decoded.ExternalID,
	}

	if err := h.changeRequestSvc.CreateChangeRequest(ctx, crq); err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	encodeResponse(ctx, writer, http.StatusCreated, newCreateChangeRequestResponse(crq))
}

type GetChangeRequestRequest struct {
	id cm.ID
}

func decodeGetChangeRequestRequest(_ context.Context, request *http.Request) *GetChangeRequestRequest {
	return &GetChangeRequestRequest{
		id: cm.ID(chi.URLParam(request, "changeRequestId")),
	}
}

type GetChangeRequestResponse struct {
	IsAutoClose bool   `json:"isAutoClose"`
	CreatedAt   int64  `json:"createdAt"`
	UpdateAt    int64  `json:"updateAt,omitempty"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func newGetChangeRequestResponse(crq *cm.ChangeRequest) *GetChangeRequestResponse {
	resp := &GetChangeRequestResponse{
		IsAutoClose: crq.IsAutoClose,
		CreatedAt:   crq.CreatedAt.UnixMilli(),
		ID:          crq.ID.String(),
		Type:        crq.Type.String(),
		Summary:     crq.Summary,
		Description: crq.Description,
	}

	if !crq.UpdatedAt.IsZero() {
		resp.UpdateAt = crq.UpdatedAt.UnixMilli()
	}

	return resp
}

func (h *ChangeRequestHandler) handleGetChangeRequest(writer http.ResponseWriter, request *http.Request) {
	var (
		ctx     = request.Context()
		decoded = decodeGetChangeRequestRequest(ctx, request)
	)

	crq, err := h.changeRequestSvc.FindChangeRequestByID(ctx, decoded.id)
	if err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	encodeResponse(request.Context(), writer, http.StatusOK, newGetChangeRequestResponse(crq))
}
