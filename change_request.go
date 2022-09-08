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
	// ID is the change request unique identifier.
	ID ID

	// Type is the change request type.
	Type ChangeRequestType

	// Summary is the short text describes the change.
	Summary string

	// Description is the full free form text describes the change.
	Description string

	// IsAutoClose is the flag that indicates that change request
	// should be closed automatically.
	IsAutoClose bool

	// CreatedAt is the UTC time when change request was created.
	CreatedAt time.Time

	// UpdateAt is the UTC time when change request was updated.
	UpdateTime time.Time
}

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService interface {
	// CreateChangeRequest creates a new change request.
	CreateChangeRequest(ctx context.Context, crq *ChangeRequest) error
}
