package pgx

import (
	"context"
	cm "github.com/morozovcookie/change-management"
)

var _ cm.IncidentService = (*IncidentService)(nil)

// IncidentService represents a service for managing incidents.
type IncidentService struct {
	queryExecer QueryExecer

	idGenerator cm.IdentifierGenerator
	timer       Timer
}

func NewIncidentService(db QueryExecer, idGenerator cm.IdentifierGenerator) *IncidentService {
	return &IncidentService{
		queryExecer: db,

		idGenerator: idGenerator,
		timer:       new(utcTimer),
	}
}

// CreateIncident creates a new incident.
func (svc *IncidentService) CreateIncident(ctx context.Context, incident *cm.Incident) error {
	return nil
}

// FindIncidentByID returns an incident by unique identifier.
func (svc *IncidentService) FindIncidentByID(ctx context.Context, id cm.ID) (*cm.Incident, error) {
	return nil, nil
}
