package awesome_api_test

import (
	"temperatures/pkg/awesome_api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RequestAwesomeAPI_MustReturnPracaDaSe(t *testing.T) {
	location := awesome_api.RequestAwesomeAPI("01001000")
	assert.Equal(t, "01001000", location.Cep)
	assert.Equal(t, "SP", location.State)
	assert.Equal(t, "São Paulo", location.City)
	assert.NotEmpty(t, location.Latitude)
	assert.NotEmpty(t, location.Longitude)
}

func Test_RequestAwesomeAPI_MustReturnInvalidCep(t *testing.T) {
	location := awesome_api.RequestAwesomeAPI("00000000")
	assert.Empty(t, location.Cep)
	assert.Empty(t, location.State)
	assert.Empty(t, location.City)
}
