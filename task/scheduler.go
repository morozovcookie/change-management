package task

import (
	"context"
	"time"
)

type Scheduler struct {
	processor QueueProcessor
	timeout   time.Duration

	done chan struct{}
}

func NewScheduler(processor QueueProcessor, timeout time.Duration) *Scheduler {
	return &Scheduler{
		processor: processor,
		timeout:   timeout,

		done: make(chan struct{}, 1),
	}
}

func (scheduler *Scheduler) Schedule(ctx context.Context) {
	ticker := time.NewTicker(scheduler.timeout)
	defer ticker.Stop()

	_ = scheduler.processor.ProcessQueue(ctx)

	for {
		select {
		case <-ctx.Done():
			scheduler.done <- struct{}{}
			close(scheduler.done)

			return
		case <-ticker.C:
			_ = scheduler.processor.ProcessQueue(ctx)
		}
	}
}

func (scheduler *Scheduler) Done() <-chan struct{} {
	return scheduler.done
}
