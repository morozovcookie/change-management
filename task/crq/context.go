package crq

import (
	"context"

	cm "github.com/morozovcookie/change-management"
)

type Context struct {
	cm.ChangeRequestUpdater
	cm.IssueService

	instance *cm.ChangeRequest
	state    State
}

func (crq *Context) Handle(ctx context.Context) error {
	return crq.state.Handle(ctx, crq)
}

func (crq *Context) Instance() *cm.ChangeRequest {
	return crq.instance
}

func (crq *Context) ChangeState(_ context.Context, state State) {
	crq.state = state
}
