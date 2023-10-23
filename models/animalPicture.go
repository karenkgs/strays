package models

import "github.com/google/uuid"

type AnimalPicture struct {
	id         uuid.UUID
	source     string
	isApproved bool
}
