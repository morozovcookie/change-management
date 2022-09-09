package pgx

import (
	"context"
	"time"
)

type Timer interface {
	Time(ctx context.Context) time.Time
}

var _ Timer = (*utcTimer)(nil)

type utcTimer struct{}

func (utcTimer) Time(_ context.Context) time.Time {
	return time.Now().UTC()
}
