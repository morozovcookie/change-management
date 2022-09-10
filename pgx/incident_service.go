package pgx

import (
	"context"
	"database/sql"
	"fmt"
	cm "github.com/morozovcookie/change-management"
	"time"
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
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	id, createdAt := svc.idGenerator.GenerateIdentifier(ctx), svc.timer.Time(ctx)

	const query = "INSERT INTO controller.incidents (incident_id, incident_summary, incident_description, created_at)" +
		" VALUES ($1, $2, $3, $4)"

	_, err := svc.queryExecer.Exec(ctx, query, id, incident.Summary, incident.Description, createdAt.UnixMilli())
	if err != nil {
		return fmt.Errorf("create incident error: %w", err)
	}

	incident.ID, incident.CreatedAt = id, createdAt

	return nil
}

// FindIncidentByID returns an incident by unique identifier.
func (svc *IncidentService) FindIncidentByID(ctx context.Context, id cm.ID) (*cm.Incident, error) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Millisecond*100))
	defer cancel()

	const query = "SELECT incident_summary, incident_description, created_at, updated_at FROM controller.incidents " +
		"WHERE incident_id = $1"

	var (
		incident = new(cm.Incident)

		createdAt int64
		updatedAt sql.NullInt64
	)

	err := svc.queryExecer.QueryRow(ctx, query, id).Scan(&incident.Summary, &incident.Description, &createdAt,
		&updatedAt)
	if err != nil {
		return nil, fmt.Errorf("find incident by id: %w", err)
	}

	incident.ID, incident.CreatedAt = id, time.UnixMilli(createdAt)

	if updatedAt.Valid {
		incident.UpdatedAt = time.UnixMilli(updatedAt.Int64)
	}

	return incident, nil
}
