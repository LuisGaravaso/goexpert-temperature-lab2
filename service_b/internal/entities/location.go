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
	coordPattern := `^-?\d+(\.\d+)?,-?\d+(\.\d+)?$`
	digitsPattern := `^\d{8}$`

	coordRegex := regexp.MustCompile(coordPattern)
	digitsRegex := regexp.MustCompile(digitsPattern)

	switch {
	case coordRegex.MatchString(l.Id):
		l.IsValid = true
		l.Type = "Coordinates"
		l.InvalidMessage = ""
	case digitsRegex.MatchString(l.Id):
		l.IsValid = true
		l.Type = "CEP"
		l.InvalidMessage = ""
	default:
		l.IsValid = false
		l.Type = "Invalid"
		l.InvalidMessage = "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates"
	}

}
