package states

import (
	"context"

	"github.com/morozovcookie/change-management/task/crq"
)

var _ crq.State = (*ClosedRequestState)(nil)

type ClosedRequestState struct{}

func (state *ClosedRequestState) Handle(_ context.Context, _ *crq.Context) error {
	panic(crq.ErrIllegalOperation)
}
