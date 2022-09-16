package pgx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/morozovcookie/change-management/task"
	"github.com/morozovcookie/change-management/task/crq"
)

var _ task.QueueProcessor = (*ChangeRequestQueueProcessor)(nil)

// ChangeRequestQueueProcessor represents a service for processing messages from change request queue.
type ChangeRequestQueueProcessor struct {
	txBeginner TxBeginner
	batchSize  int

	timer Timer
}

func NewChangeRequestQueueProcessor(txBeginner TxBeginner, batchSize int) *ChangeRequestQueueProcessor {
	return &ChangeRequestQueueProcessor{
		txBeginner: txBeginner,
		batchSize:  batchSize,

		timer: new(utcTimer),
	}
}

type changeRequestQueueItem struct {
	id        int64
	content   json.RawMessage
	failCount int16

	context *crq.Context
}

type processQueueItemResult struct {
	id        int64
	failCount int64
	err       error
}

// ProcessQueue process messages from queue.
func (svc *ChangeRequestQueueProcessor) ProcessQueue(ctx context.Context) error {
	tx, err := svc.txBeginner.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("process queue: %w", err)
	}

	defer func(ctx context.Context, tx pgx.Tx, err *error) {
		if err == nil || *err == nil {
			return
		}

		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			*err = rollbackErr
		}
	}(ctx, tx, &err)

	query := "SELECT row_id, content, fail_count FROM controller.change_request_queue ORDER BY row_id FOR UPDATE" +
		" SKIP LOCKED LIMIT $1"

	rows, err := tx.Query(ctx, query, svc.batchSize)
	if err != nil {
		return fmt.Errorf("process queue: %w", err)
	}

	defer rows.Close()

	ii := make([]*changeRequestQueueItem, 0, svc.batchSize)

	for rows.Next() {
		item := new(changeRequestQueueItem)

		if err := rows.Scan(&item.id, &item.content, &item.failCount); err != nil {
			return fmt.Errorf("process queue: %w", err)
		}

		ii = append(ii)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("process queue: %w", err)
	}

	for i, item := range ii {
		ii[i].context = queueItemContentToChangeRequestContext(item.content)
	}

	resultCh := make(chan *processQueueItemResult, 1)
	defer close(resultCh)

	for _, item := range ii {
		go func(ctx context.Context, item *changeRequestQueueItem, out chan<- *processQueueItemResult) {
			result := &processQueueItemResult{
				id: item.id,
			}

			if result.err = item.context.Handle(ctx); result.err != nil {
				result.failCount++
			}

			out <- result
		}(ctx, item, resultCh)
	}

	forDelete := make([]int64, 0, svc.batchSize)

	query = "UPDATE controller.change_request_queue SET fail_count = $1, last_error = $2, updated_at = $3 WHERE " +
		"row_id = $4"

	for range ii {
		result := <-resultCh

		if err := result.err; err == nil {
			forDelete = append(forDelete, result.id)

			continue
		}

		_, _ = tx.Exec(ctx, query, result.failCount, result.err.Error(), svc.timer.Time(ctx).UnixMilli(), result.id)
	}

	if err = deleteProcessedItems(ctx, tx, forDelete); err != nil {
		return fmt.Errorf("process queue: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("process queue: %w", err)
	}

	return nil
}

func queueItemContentToChangeRequestContext(_ json.RawMessage) *crq.Context {
	return &crq.Context{}
}

func deleteProcessedItems(ctx context.Context, tx pgx.Tx, ii []int64) error {
	if len(ii) == 0 {
		return nil
	}

	var (
		buf    = bytes.NewBufferString("DELETE FROM controller.change_request_queue WHERE row_id IN (")
		suffix = ", "

		args = make([]any, len(ii))
	)

	for i := range ii {
		args[i] = ii[i]

		if i == len(ii)-1 {
			suffix = ")"
		}

		_, _ = fmt.Fprintf(buf, "$%d%s", i+1, suffix)
	}

	if _, err := tx.Exec(ctx, buf.String(), args...); err != nil {
		return err
	}

	return nil
}
