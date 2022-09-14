package v1

import (
	"context"

	cm "github.com/morozovcookie/change-management"
)

// IssueService represents a service for managing issues.
type IssueService struct{}

// CreateIssue creates a new ticket in issue tracker system.
func (svc *IssueService) CreateIssue(_ context.Context, issue *cm.Issue) error {
	issue.Key = "KEY-123"

	return nil
}

// CloseIssue closes a new ticket in issue tracker system.
func (svc *IssueService) CloseIssue(_ context.Context, _ string) error {
	return nil
}
