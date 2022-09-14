package states

import (
	"context"

	"github.com/morozovcookie/change-management/task/crq"
)

var _ crq.State = (*RegisteredRequestState)(nil)

type RegisteredRequestState struct {
	container crq.StateContainer
}

func (state *RegisteredRequestState) Handle(ctx context.Context, crq *crq.Context) error {
	// TODO: close issue in issue tracker

	crq.ChangeState(ctx, state.container.ClosedRequestState(ctx))

	return nil
}
