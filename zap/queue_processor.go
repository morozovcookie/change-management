package zap

import (
	"context"
	"time"

	"github.com/morozovcookie/change-management/task"
	"go.uber.org/zap"
)

var _ task.QueueProcessor = (*QueueProcessor)(nil)

// QueueProcessor represents a service for processing messages from queue.
type QueueProcessor struct {
	wrapped task.QueueProcessor
	logger  *zap.Logger
}

func NewQueueProcessor(processor task.QueueProcessor, logger *zap.Logger) *QueueProcessor {
	return &QueueProcessor{
		wrapped: processor,
		logger:  logger,
	}
}

// ProcessQueue process messages from queue.
func (svc *QueueProcessor) ProcessQueue(ctx context.Context) error {
	start := time.Now().UTC()

	err := svc.wrapped.ProcessQueue(ctx)

	end := time.Now().UTC()

	ff := []zap.Field{
		zap.Error(err), zap.Stringer("start", start), zap.Stringer("end", end),
		zap.Stringer("elapsed", end.Sub(start).Round(time.Millisecond)),
	}

	svc.logger.Debug("process queue", ff...)

	if err != nil {
		svc.logger.Error("process queue", ff...)

		return err //nolint:wrapcheck
	}

	return nil
}
