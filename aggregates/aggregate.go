package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/eventsourcing"
	"github.com/google/uuid"
)

type AggregateType string

type Aggregate struct {
	id            uuid.UUID
	aggregateType AggregateType
	version       int
	events        []eventsourcing.Event
}
