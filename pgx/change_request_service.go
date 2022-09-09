package pgx

import (
	"context"
	"fmt"
	"time"

	cm "github.com/morozovcookie/change-management"
)

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService struct {
	db QueryExecer

	idGenerator cm.IdentifierGenerator
	timer       Timer
}

func NewChangeRequestService(db QueryExecer, idGenerator cm.IdentifierGenerator) *ChangeRequestService {
	return &ChangeRequestService{
		db: db,

		idGenerator: idGenerator,
		timer:       new(utcTimer),
	}
}

// CreateChangeRequest creates a new change request.
func (svc *ChangeRequestService) CreateChangeRequest(ctx context.Context, crq *cm.ChangeRequest) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	id, createdAt := svc.idGenerator.GenerateIdentifier(ctx), svc.timer.Time(ctx)

	const query = "INSERT INTO change_requests (crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, " +
		"created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := svc.db.Exec(ctx, query, id, crq.Type.String(), crq.Summary, crq.Description, crq.IsAutoClose, createdAt)
	if err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	crq.ID, crq.CreatedAt = id, createdAt

	return nil
}

// FindChangeRequestByID returns change request by unique identifier.
func (svc *ChangeRequestService) FindChangeRequestByID(ctx context.Context, id cm.ID) (*cm.ChangeRequest, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	const query = "SELECT crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, created_at, updated_at " +
		"FROM change_requests WHERE crq_id = $1"

	var (
		crq = new(cm.ChangeRequest)

		requestType string
		createdAt   int64
		updatedAt   int64
	)

	err := svc.db.QueryRow(ctx, query, id.String()).Scan(&crq.ID, &requestType, &crq.Summary, &crq.Description,
		&crq.IsAutoClose, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("find change request by id: %w", err)
	}

	if crq.Type = cm.ChangeRequestType(requestType); !crq.Type.IsValid() {
		return nil, fmt.Errorf("unsupported change request type: %s", requestType)
	}

	crq.CreatedAt, crq.UpdatedAt = time.UnixMilli(createdAt), time.UnixMilli(updatedAt)

	return crq, nil
}
