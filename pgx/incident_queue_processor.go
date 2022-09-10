package pgx

import (
	"context"

	"github.com/morozovcookie/change-management/task"
)

var _ task.QueueProcessor = (*IncidentQueueProcessor)(nil)

// IncidentQueueProcessor represents a service for processing messages from incident queue.
type IncidentQueueProcessor struct {
	txBeginner TxBeginner
}

func NewIncidentQueueProcessor(txBeginner TxBeginner) *IncidentQueueProcessor {
	return &IncidentQueueProcessor{
		txBeginner: txBeginner,
	}
}

// ProcessQueue process messages from queue.
func (svc *IncidentQueueProcessor) ProcessQueue(_ context.Context) error {
	return nil
}
