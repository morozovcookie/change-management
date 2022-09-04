package cm

import (
	"time"
)

// ChangeRequest is the request for making the change of current infrastructure
// state.
type ChangeRequest struct {
	// ID is the change request unique identifier.
	ID ID

	// CreatedAt is the UTC time when change request was created.
	CreatedAt time.Time

	// UpdateAt is the UTC time when change request was updated.
	UpdateTime time.Time
}

// ChangeRequestService represents a service for manging ChangeRequest data.
type ChangeRequestService interface{}
