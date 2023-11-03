package eventsourcing

import "github.com/google/uuid"

type AggregateType string

type Aggregate struct {
	id            uuid.UUID
	aggregateType AggregateType
	version       int
	events        []Event
}
