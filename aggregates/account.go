package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/models"
	"github.com/google/uuid"
)

type Account struct {
	id          uuid.UUID
	firstName   string
	middleName  string
	lastName    string
	phoneNumber models.PhoneNumber
	isVet       bool
	country     models.Country
	state       string
}
