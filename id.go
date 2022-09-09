package cm

import (
	"context"
	"fmt"
)

var _ fmt.Stringer = (*ID)(nil)

type ID string

func (id ID) String() string {
	return string(id)
}

const EmptyID = ID("")

// IdentifierGenerator represents a service for generating unique identifier
// value.
type IdentifierGenerator interface {
	// GenerateIdentifier returns a unique identifier value.
	GenerateIdentifier(ctx context.Context) ID
}
