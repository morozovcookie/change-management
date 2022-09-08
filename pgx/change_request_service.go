package pgx

import (
	"context"

	cm "github.com/morozovcookie/change-management"
)

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService struct{}

func NewChangeRequestService() *ChangeRequestService {
	return &ChangeRequestService{}
}

// CreateChangeRequest creates a new change request.
func (svc *ChangeRequestService) CreateChangeRequest(ctx context.Context, crq *cm.ChangeRequest) error {
	return nil
}
