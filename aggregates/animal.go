package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/models"
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	Aggregate
	accountID   uuid.UUID
	createdAt   time.Time
	firstSeenAt models.Address
	isDesexed   bool
	conditions  []models.Condition
	picture     models.Picture
}
