package entities

import "regexp"

type Location struct {
	Id             string
	Type           string
	IsValid        bool
	InvalidMessage string
}

func NewLocation(id string) *Location {

	newLocation := &Location{Id: id}
	newLocation.Validate()
	return newLocation
}

func (l *Location) Validate() {
	digitsPattern := `^\d{8}$`

	digitsRegex := regexp.MustCompile(digitsPattern)

	switch {
	case digitsRegex.MatchString(l.Id):
		l.IsValid = true
		l.Type = "CEP"
		l.InvalidMessage = ""
	default:
		l.IsValid = false
		l.Type = "Invalid"
		l.InvalidMessage = "Must be in the format 01001001 for CEP"
	}

}
