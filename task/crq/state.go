package crq

import (
	"context"
	"errors"
)

var ErrIllegalOperation = errors.New("operation is illegal")

type State interface {
	Handle(ctx context.Context, crq *Context) error
}

type StateContainer interface {
	NewRequestState(ctx context.Context) State
	RegisteredRequestState(ctx context.Context) State
	ClosedRequestState(ctx context.Context) State
}
