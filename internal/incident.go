package internal

import (
	"time"
)

// Incident represents a user appeal and service request.
type Incident struct {
	// ID is the incident unique identifier.
	ID ID

	// Summary is the short-form description of incident.
	Summary string

	// Description is the long-form description of incident.
	// Description contains detailed and human-readable information what
	// incident is about.
	Description string

	// CreatedAt is the time when incident was created.
	// CreatedAt MUST be always proccessed in UTC location
	// and be converted to user location on the client side.
	CreatedAt time.Time

	// UpdatedAt is the time when incident was updated at the last time.
	// UpdatedAt could have a zero-value in case if incident was not updated
	// yet.
	// UpdatedAt MUST be always processed in UTC location and be converted to
	// user location on the client side.
	UpdatedAt time.Time
}

// IncidentService represents a service for managing incidents.
type IncidentService interface{}
