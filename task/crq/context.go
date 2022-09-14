package crq

import (
	"context"
)

type Context struct {
	state State
}

func (crq *Context) Handle(ctx context.Context) error {
	return crq.state.Handle(ctx, crq)
}

func (crq *Context) ChangeState(_ context.Context, state State) {
	crq.state = state
}
