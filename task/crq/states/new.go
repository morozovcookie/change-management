package states

import (
	"context"

	"github.com/morozovcookie/change-management/task/crq"
)

var _ crq.State = (*NewRequestState)(nil)

type NewRequestState struct {
	container crq.StateContainer
}

func (state *NewRequestState) Handle(ctx context.Context, crq *crq.Context) error {
	//	TODO: register CRQ in issue tracker

	crq.ChangeState(ctx, state.container.RegisteredRequestState(ctx))

	return nil
}
