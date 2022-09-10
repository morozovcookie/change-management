package task

import (
	"context"
)

// QueueProcessor represents a service for processing messages from queue.
type QueueProcessor interface {
	// ProcessQueue process messages from queue.
	ProcessQueue(ctx context.Context) error
}
