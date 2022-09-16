package pgx

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"

	cm "github.com/morozovcookie/change-management"
)

var _ cm.ChangeRequestService = (*ChangeRequestService)(nil)

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService struct {
	conn Conn

	idGenerator cm.IdentifierGenerator
	timer       Timer
}

func NewChangeRequestService(conn Conn, idGenerator cm.IdentifierGenerator) *ChangeRequestService {
	return &ChangeRequestService{
		conn: conn,

		idGenerator: idGenerator,
		timer:       new(utcTimer),
	}
}

// CreateChangeRequest creates a new change request.
func (svc *ChangeRequestService) CreateChangeRequest(ctx context.Context, crq *cm.ChangeRequest) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	query := "SELECT crq_id, created_at FROM controller.change_requests WHERE hash = $1 ORDER BY row_id DESC LIMIT 1"

	err := svc.conn.QueryRow(ctx, query, calculateChangeRequestHash(crq)).Scan(&crq.ID, &crq.CreatedAt)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("create change request: %w", err)
	}

	if err == nil {
		return nil
	}

	tx, err := svc.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	defer func(ctx context.Context, tx pgx.Tx, err *error) {
		if err == nil || *err == nil {
			return
		}

		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			*err = rollbackErr
		}
	}(ctx, tx, &err)

	id, createdAt := svc.idGenerator.GenerateIdentifier(ctx), svc.timer.Time(ctx)

	query = "INSERT INTO controller.change_requests (crq_id, crq_type, crq_summary, crq_description, " +
		"crq_is_auto_close, created_at) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err = tx.Exec(ctx, query, id, crq.Type.String(), crq.Summary, crq.Description, crq.IsAutoClose,
		createdAt.UnixMilli())
	if err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	query = "INSERT INTO controller.change_request_queue (content, created_at) VALUES ($1, $2)"

	if _, err := tx.Exec(ctx, query, json.RawMessage(nil), svc.timer.Time(ctx).UnixMilli()); err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	crq.ID, crq.CreatedAt = id, createdAt

	return nil
}

func calculateChangeRequestHash(crq *cm.ChangeRequest) string {
	hash := sha256.New()

	_, _ = fmt.Fprintf(hash, "%t%s%s%s", crq.IsAutoClose, crq.Type, crq.Summary, crq.Description)

	return hex.EncodeToString(hash.Sum(nil))
}

// FindChangeRequestByID returns change request by unique identifier.
func (svc *ChangeRequestService) FindChangeRequestByID(ctx context.Context, id cm.ID) (*cm.ChangeRequest, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	const query = "SELECT crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, created_at, updated_at " +
		"FROM controller.change_requests WHERE crq_id = $1"

	var (
		crq = new(cm.ChangeRequest)

		requestType string
		createdAt   int64
		updatedAt   sql.NullInt64
	)

	err := svc.db.QueryRow(ctx, query, id.String()).Scan(&crq.ID, &requestType, &crq.Summary, &crq.Description,
		&crq.IsAutoClose, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("find change request by id: %w", err)
	}

	if crq.Type = cm.ChangeRequestType(requestType); !crq.Type.IsValid() {
		return nil, fmt.Errorf("unsupported change request type: %s", requestType)
	}

	crq.CreatedAt = time.UnixMilli(createdAt)

	if updatedAt.Valid {
		crq.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	return crq, nil
}

// UpdateChangeRequest updates an existent change request.
func (svc *ChangeRequestService) UpdateChangeRequest(_ context.Context, _ *cm.ChangeRequest) error {
	return nil
}
