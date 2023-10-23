package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/models"
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	id          uuid.UUID
	accountID   uuid.UUID
	createdAt   time.Time
	firstSeenAt models.Address
	isDesexed   bool
	conditions  []models.Condition
	picture     models.AnimalPicture
}
