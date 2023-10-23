package models

import (
	"github.com/google/uuid"
	"time"
)

type Condition struct {
	id                   uuid.UUID
	name                 string
	requiresVetAttention bool
	lastTreatedAt        time.Time
}
