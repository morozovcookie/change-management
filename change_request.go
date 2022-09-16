package cm

import (
	"context"
	"fmt"
	"time"
)

var _ fmt.Stringer = (*ChangeRequestType)(nil)

// ChangeRequestType represents the type of change request.
type ChangeRequestType string

const (
	ChangeRequestTypeCRQ     = "CRQ"
	ChangeRequestTypeAutoCRQ = "AutoCRQ"
)

func (t ChangeRequestType) String() string {
	return string(t)
}

func (t ChangeRequestType) IsValid() bool {
	switch t {
	case ChangeRequestTypeCRQ, ChangeRequestTypeAutoCRQ:
		return true
	default:
	}

	return false
}

// ChangeRequest is the request for making the change of current infrastructure
// state.
type ChangeRequest struct {
	// IsAutoClose is the flag that indicates that change request
	// should be closed automatically.
	IsAutoClose bool

	// Issue is the linked ticker in issue tracker system.
	Issue *Issue

	// ID is the change request unique identifier.
	ID ID

	// ExternalID is the identifier in the API consumer system.
	ExternalID string

	// Type is the change request type.
	Type ChangeRequestType

	// Summary is the short text describes the change.
	Summary string

	// Description is the full free form text describes the change.
	Description string

	// CreatedAt is the UTC time when change request was created.
	CreatedAt time.Time

	// UpdatedAt is the UTC time when change request was updated.
	UpdatedAt time.Time
}

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService interface {
	// CreateChangeRequest creates a new change request.
	CreateChangeRequest(ctx context.Context, crq *ChangeRequest) error

	// FindChangeRequestByID returns change request by unique identifier.
	FindChangeRequestByID(ctx context.Context, id ID) (*ChangeRequest, error)

	// FindChangeRequestByExternalID returns change request by external
	// identifier.
	FindChangeRequestByExternalID(ctx context.Context, id string) (*ChangeRequest, error)
}

// ChangeRequestUpdater represents a service for updating change request.
type ChangeRequestUpdater interface {
	// UpdateChangeRequest updates an existent change request.
	UpdateChangeRequest(ctx context.Context, crq *ChangeRequest) error
}
