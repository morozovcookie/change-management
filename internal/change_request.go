package internal

import (
	"time"
)

// ChangeRequestType represents an enumeration of possible change request types.
type ChangeRequestType uint8

const (
	// ChangeRequestTypeCRQ is the type of change request that requires
	// an approval from supervisor users.
	ChangeRequestTypeCRQ ChangeRequestType = iota + 1

	// ChangeRequestTypeAutoCRQ is the type of change request that not require
	// any approval from supervisor users.
	// ChangeRequest with ChangeRequestTypeAutoCRQ type could be closed manually
	// or automatically.
	ChangeRequestTypeAutoCRQ
)

// ChangeRequest is the request for making the change of current infrastructure
// state.
type ChangeRequest struct {
	// ID is the change request unique identifier.
	ID ID

	// Summary is the short-form description of change.
	Summary string

	// Description is the long-form description of change.
	// Description contains detailed and human-readable information about what
	// was or will be changed.
	Description string

	// Type is the type of change request.
	Type ChangeRequestType

	// IsAutoClose is the flag that indicates that this change request need
	// to be closed automatically.
	// IsAutoClosed MUST be used only with change requests that has type of
	// ChangeRequestTypeAutoCRQ.
	IsAutoClose bool

	// CreatedAt is the time when change request was created.
	// CreatedAt MUST be always proccessed in UTC location
	// and be converted to user location on the client side.
	CreatedAt time.Time

	// UpdatedAt is the time when change request was updated at the last time.
	// UpdatedAt could have a zero-value in case if change request was not
	// updated yet.
	// UpdatedAt MUST be always processed in UTC location and be converted to
	// user location on the client side.
	UpdatedAt time.Time
}

// ChangeRequestService represents a service for managing change requests.
type ChangeRequestService interface{}
