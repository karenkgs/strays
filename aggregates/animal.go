package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/eventsourcing"
	"Users/karen/Developer/Personal/go/strays/models"
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	eventsourcing.Aggregate
	accountID   uuid.UUID
	createdAt   time.Time
	firstSeenAt models.Address
	isDesexed   bool
	conditions  []models.Condition
	picture     models.Picture
}
