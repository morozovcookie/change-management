package cm

import (
	"context"
	"time"
)

// Incident represents a user appeal and service request.
type Incident struct {
	// ID is the incident unique identifier.
	ID ID

	// Summary is the short text describes the change.
	Summary string

	// Description is the full free form text describes the change.
	Description string

	// CreatedAt is the UTC time when incident was created.
	CreatedAt time.Time

	// UpdatedAt is the UTC time when incident was updated.
	UpdatedAt time.Time
}

// IncidentService represents a service for managing incidents.
type IncidentService interface {
	// CreateIncident creates a new incident.
	CreateIncident(ctx context.Context, incident *Incident) error

	// FindIncidentByID returns an incident by unique identifier.
	FindIncidentByID(ctx context.Context, id ID) (*Incident, error)
}
