package states

import (
	"context"
	"fmt"

	cm "github.com/morozovcookie/change-management"
	"github.com/morozovcookie/change-management/task/crq"
)

var _ crq.State = (*NewRequestState)(nil)

type NewRequestState struct {
	container crq.StateContainer
}

func (state *NewRequestState) Handle(ctx context.Context, crq *crq.Context) error {
	if err := state.createIssue(ctx, crq); err != nil {
		return fmt.Errorf("handle New request state: %w", err)
	}

	if err := state.updateChangeRequest(ctx, crq); err != nil {
		return fmt.Errorf("handle New request state: %w", err)
	}

	crq.ChangeState(ctx, state.container.RegisteredRequestState(ctx))

	return nil
}

func (state *NewRequestState) createIssue(ctx context.Context, crq *crq.Context) error {
	issue := &cm.Issue{}

	if err := crq.IssueService.CreateIssue(ctx, issue); err != nil {
		return err //nolint:wrapcheck
	}

	crq.Instance().Issue = issue

	return nil
}

func (state *NewRequestState) updateChangeRequest(ctx context.Context, crq *crq.Context) error {
	return crq.ChangeRequestUpdater.UpdateChangeRequest(ctx, crq.Instance()) //nolint:wrapcheck
}
