package events

import (
	"Users/karen/Developer/Personal/go/strays/eventsourcing"
	"Users/karen/Developer/Personal/go/strays/models"
)

type AccountCreated struct {
	eventsourcing.Event
	firstName       string
	middleName      string
	lastName        string
	phoneNumber     models.PhoneNumber
	isVet           bool
	canVerifyAnimal bool
	country         models.Country
	state           string
}
