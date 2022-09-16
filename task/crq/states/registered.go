package states

import (
	"context"
	"fmt"

	"github.com/morozovcookie/change-management/task/crq"
)

var _ crq.State = (*RegisteredRequestState)(nil)

type RegisteredRequestState struct {
	container crq.StateContainer
}

func (state *RegisteredRequestState) Handle(ctx context.Context, crq *crq.Context) error {
	if err := state.closeIssue(ctx, crq); err != nil {
		return fmt.Errorf("handle Registered request state: %w", err)
	}

	if err := state.updateChangeRequest(ctx, crq); err != nil {
		return fmt.Errorf("handle Registered request state: %w", err)
	}

	crq.ChangeState(ctx, state.container.ClosedRequestState(ctx))

	return nil
}

func (state *RegisteredRequestState) closeIssue(ctx context.Context, crq *crq.Context) error {
	return crq.IssueService.CloseIssue(ctx, crq.Instance().Issue.Key) //nolint:wrapcheck
}

func (state *RegisteredRequestState) updateChangeRequest(ctx context.Context, crq *crq.Context) error {
	return crq.ChangeRequestUpdater.UpdateChangeRequest(ctx, crq.Instance()) //nolint:wrapcheck
}
