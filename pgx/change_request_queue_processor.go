package pgx

import (
	"context"

	"github.com/morozovcookie/change-management/task"
)

var _ task.QueueProcessor = (*ChangeRequestQueueProcessor)(nil)

// ChangeRequestQueueProcessor represents a service for processing messages from change request queue.
type ChangeRequestQueueProcessor struct {
	txBeginner TxBeginner
}

func NewChangeRequestQueueProcessor(txBeginner TxBeginner) *ChangeRequestQueueProcessor {
	return &ChangeRequestQueueProcessor{
		txBeginner: txBeginner,
	}
}

// ProcessQueue process messages from queue.
func (svc *ChangeRequestQueueProcessor) ProcessQueue(_ context.Context) error {
	return nil
}
