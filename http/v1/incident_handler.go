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
	IncidentHandlerPathPrefix = "/api/v1/incidents"
	CreateIncidentPathPrefix  = "/"
	GetIncidentPathPrefix     = "/{incidentId}"
)

var _ http.Handler = (*IncidentHandler)(nil)

type IncidentHandler struct {
	http.Handler

	incidentSvc cm.IncidentService
}

func NewIncidentHandler(incidentSvc cm.IncidentService) *IncidentHandler {
	var (
		router  = chi.NewRouter()
		handler = &IncidentHandler{
			Handler: router,

			incidentSvc: incidentSvc,
		}
	)

	router.Post(CreateIncidentPathPrefix, handler.handleCreateIncident)
	router.Get(GetIncidentPathPrefix, handler.handleGetIncident)

	return handler
}

type CreateIncidentRequest struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func decodeCreateIncidentRequest(_ context.Context, reader io.Reader) (*CreateIncidentRequest, error) {
	decoded := new(CreateIncidentRequest)

	if err := json.NewDecoder(reader).Decode(decoded); err != nil {
		return nil, &cm.Error{
			Code:    cm.ErrorCodeInvalid,
			Message: "failed to decode CreateIncident request",
			Err:     err,
		}
	}

	return decoded, nil
}

type CreateIncidentResponse struct {
	ID string `json:"id"`
}

func newCreateIncidentResponse(incident *cm.Incident) *CreateIncidentResponse {
	return &CreateIncidentResponse{
		ID: incident.ID.String(),
	}
}

func (h *IncidentHandler) handleCreateIncident(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	decoded, err := decodeCreateIncidentRequest(ctx, request.Body)
	if err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	incident := &cm.Incident{
		Summary:     decoded.Summary,
		Description: decoded.Description,
	}

	if err := h.incidentSvc.CreateIncident(ctx, incident); err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	encodeResponse(ctx, writer, http.StatusOK, newCreateIncidentResponse(incident))
}

type GetIncidentRequest struct {
	ID cm.ID
}

func decodeGetIncidentRequest(_ context.Context, request *http.Request) *GetIncidentRequest {
	return &GetIncidentRequest{
		ID: cm.ID(chi.URLParam(request, "incidentId")),
	}
}

type GetIncidentResponse struct {
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt,omitempty"`
	ID          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func newGetIncidentResponse(incident *cm.Incident) *GetIncidentResponse {
	resp := &GetIncidentResponse{
		CreatedAt:   incident.CreatedAt.UnixMilli(),
		ID:          incident.ID.String(),
		Summary:     incident.Summary,
		Description: incident.Description,
	}

	if !incident.UpdatedAt.IsZero() {
		resp.UpdatedAt = incident.UpdatedAt.UnixMilli()
	}

	return resp
}

func (h *IncidentHandler) handleGetIncident(writer http.ResponseWriter, request *http.Request) {
	var (
		ctx     = request.Context()
		decoded = decodeGetIncidentRequest(ctx, request)
	)

	incident, err := h.incidentSvc.FindIncidentByID(ctx, decoded.ID)
	if err != nil {
		encodeErrorResponse(ctx, writer, err)

		return
	}

	encodeResponse(ctx, writer, http.StatusOK, newGetIncidentResponse(incident))
}
