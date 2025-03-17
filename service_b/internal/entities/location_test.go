package entities_test

import (
	"temperatures/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	location := entities.NewLocation("01001001")
	assert.Equal(t, location.Id, "01001001")
}

func TestLocationType(t *testing.T) {

	// Desired Formats
	location := entities.NewLocation("-21.5,64.5")
	assert.Equal(t, location.IsValid, true)
	assert.Equal(t, location.Type, "Coordinates")

	location = entities.NewLocation("01001001")
	assert.Equal(t, location.IsValid, true)
	assert.Equal(t, location.Type, "CEP")

	// Invalid Formats
	location = entities.NewLocation("")
	assert.Equal(t, location.IsValid, false)
	assert.Equal(t, location.Type, "Invalid")
	assert.Equal(t, location.InvalidMessage, "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates")

	location = entities.NewLocation("010010-01")
	assert.Equal(t, location.IsValid, false)
	assert.Equal(t, location.Type, "Invalid")
	assert.Equal(t, location.InvalidMessage, "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates")

	location = entities.NewLocation("1")
	assert.Equal(t, location.IsValid, false)
	assert.Equal(t, location.Type, "Invalid")
	assert.Equal(t, location.InvalidMessage, "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates")

	location = entities.NewLocation("123456789")
	assert.Equal(t, location.IsValid, false)
	assert.Equal(t, location.Type, "Invalid")
	assert.Equal(t, location.InvalidMessage, "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates")
}
