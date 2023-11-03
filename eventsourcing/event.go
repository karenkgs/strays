package eventsourcing

import (
	"Users/karen/Developer/Personal/go/strays/aggregates"
	"github.com/google/uuid"
	"time"
)

type EventType string

type Event struct {
	id            uuid.UUID
	aggregateType aggregates.AggregateType
	aggregateID   uuid.UUID
	position      int
	eventType     EventType
	Payload       interface{}
	RecordedAt    time.Time
}
