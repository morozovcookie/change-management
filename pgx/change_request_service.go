package pgx

import (
	"context"
	"database/sql"
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

const (
	createChangeRequestRecordQuery = `INSERT INTO controller.change_requests 
(crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, crq_external_id, created_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7)`

	putItemIntoChangeRequestQueueQuery = `INSERT INTO controller.change_request_queue 
(content, created_at) 
VALUES ($1, $2)`
)

// CreateChangeRequest creates a new change request.
func (svc *ChangeRequestService) CreateChangeRequest(ctx context.Context, crq *cm.ChangeRequest) error {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	err := svc.checkChangeRequestExistence(ctx, crq.ExternalID)
	if err != nil && cm.ErrorCodeFromError(err) != cm.ErrorCodeNotFound {
		return err
	}

	if err == nil {
		return fmt.Errorf("create change reqeust: %w", &cm.Error{
			Code:    cm.ErrorCodeConflict,
			Message: fmt.Sprintf(`change request with external identifier "%s" already exist`, crq.ExternalID),
			Err:     nil,
		})
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

	_, err = tx.Exec(ctx, createChangeRequestRecordQuery, id, crq.Type.String(), crq.Summary, crq.Description,
		crq.IsAutoClose, crq.ExternalID, createdAt.UnixMilli())
	if err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	_, err = tx.Exec(ctx, putItemIntoChangeRequestQueueQuery, json.RawMessage(nil), svc.timer.Time(ctx).UnixMilli())
	if err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("create change request: %w", err)
	}

	crq.ID, crq.CreatedAt = id, createdAt

	return nil
}

const checkChangeRequestExistenceQuery = `SELECT exists(
    SELECT 1 
    FROM controller.change_requests 
    WHERE crq_external_id = $1
)`

func (svc *ChangeRequestService) checkChangeRequestExistence(ctx context.Context, id string) error {
	var isExist bool

	if err := svc.conn.QueryRow(ctx, checkChangeRequestExistenceQuery, id).Scan(&isExist); err != nil {
		return err
	}

	if !isExist {
		return &cm.Error{
			Code:    cm.ErrorCodeNotFound,
			Message: fmt.Sprintf(`change request with external identifier "%s" does not exist`, id),
			Err:     nil,
		}
	}

	return nil
}

const findChangeRequestByIdQuery = `SELECT crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, 
       crq_external_id, created_at, updated_at 
FROM controller.change_requests 
WHERE crq_id = $1`

// FindChangeRequestByID returns change request by unique identifier.
func (svc *ChangeRequestService) FindChangeRequestByID(ctx context.Context, id cm.ID) (*cm.ChangeRequest, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	var (
		crq = new(cm.ChangeRequest)

		requestType string
		createdAt   int64
		updatedAt   sql.NullInt64
	)

	err := svc.conn.QueryRow(ctx, findChangeRequestByIdQuery, id.String()).Scan(&crq.ID, &requestType, &crq.Summary,
		&crq.Description, &crq.IsAutoClose, &crq.ExternalID, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("find change request by id: %w", err)
	}

	if crq.Type = cm.ChangeRequestType(requestType); !crq.Type.IsValid() {
		return nil, fmt.Errorf("unsupported change request type: %s", requestType)
	}

	crq.CreatedAt = time.UnixMilli(createdAt)

	if val, ok := updatedAt.Int64, updatedAt.Valid; ok {
		crq.UpdatedAt = time.UnixMilli(val)
	}

	return crq, nil
}

const findChangeRequestByExternalIdQuery = `SELECT crq_id, crq_type, crq_summary, crq_description, crq_is_auto_close, 
       created_at, updated_at 
FROM controller.change_requests 
WHERE crq_external_id = $1 
ORDER BY row_id DESC 
LIMIT 1`

// FindChangeRequestByExternalID returns change request by external
// identifier.
func (svc *ChangeRequestService) FindChangeRequestByExternalID(
	ctx context.Context,
	id string,
) (
	*cm.ChangeRequest,
	error,
) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	var (
		crq = new(cm.ChangeRequest)

		createdAt int64
		updatedAt sql.NullInt64
	)

	err := svc.conn.QueryRow(ctx, findChangeRequestByExternalIdQuery, id).Scan(&crq.ID, &crq.Type, &crq.Summary,
		&crq.Description, &crq.IsAutoClose, &createdAt, &updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("find change request by external identifier: %w", &cm.Error{
			Code:    cm.ErrorCodeNotFound,
			Message: fmt.Sprintf(`change request with externa identifier "%s" does not exist`, id),
			Err:     err,
		})
	}

	if err != nil {
		return nil, fmt.Errorf("find change request by external identifier: %w", err)
	}

	crq.ExternalID, crq.CreatedAt = id, time.UnixMilli(createdAt)

	if val, ok := updatedAt.Int64, updatedAt.Valid; ok {
		crq.UpdatedAt = time.UnixMilli(val)
	}

	return crq, nil
}
