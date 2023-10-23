package models

type Address struct {
	line1    string
	line2    string
	city     string
	state    string
	postcode int16
	country  Country
}
