package aggregates

import (
	"Users/karen/Developer/Personal/go/strays/eventsourcing"
	"Users/karen/Developer/Personal/go/strays/models"
)

type Account struct {
	eventsourcing.Aggregate
	firstName       string
	middleName      string
	lastName        string
	phoneNumber     models.PhoneNumber
	isVet           bool
	canVerifyAnimal bool
	country         models.Country
	state           string
}
