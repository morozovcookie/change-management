package task

import (
	"context"
	"time"
)

type Scheduler struct {
	processor QueueProcessor
	timeout   time.Duration
}

func NewScheduler(processor QueueProcessor, timeout time.Duration) *Scheduler {
	return &Scheduler{
		processor: processor,
		timeout:   timeout,
	}
}

func (s *Scheduler) Schedule(ctx context.Context) <-chan struct{} {
	var (
		ticker = time.NewTicker(s.timeout)
		quit   = make(chan struct{}, 1)
	)

	go func(ctx context.Context, ticker *time.Ticker, processor QueueProcessor, quit chan<- struct{}) {
		defer close(quit)
		defer ticker.Stop()

		_ = processor.ProcessQueue(ctx)

		for {
			select {
			case <-ctx.Done():
				quit <- struct{}{}

				return
			case <-ticker.C:
				_ = processor.ProcessQueue(ctx)
			}
		}
	}(ctx, ticker, s.processor, quit)

	return quit
}
