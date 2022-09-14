package cm

import (
	"context"
)

// Issue represents a ticket in issue tracker system.
type Issue struct {
	Key string
}

// IssueService represents a service for managing issues.
type IssueService interface {
	// CreateIssue creates a new ticket in issue tracker system.
	CreateIssue(ctx context.Context, issue *Issue) error

	// CloseIssue closes a new ticket in issue tracker system.
	CloseIssue(ctx context.Context, key string) error
}
