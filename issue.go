package cm

import (
	"context"
)

type Issue struct{}

type IssueService interface {
	CreateIssue(ctx context.Context, issue *Issue) error
}
