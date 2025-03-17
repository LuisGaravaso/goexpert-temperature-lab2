package location_test

import (
	"log"
	"testing"

	"temperatures/internal/gateway/location"

	"github.com/stretchr/testify/assert"
)

func Test_RequestAwesomeAPI_MustConvertCEP(t *testing.T) {
	gateway := &location.AwesomeAPILocationGateway{}
	coordinates, err := gateway.Cep2Coordinates("01001000")
	assert.Nil(t, err)
	assert.NotEmpty(t, coordinates)

	coordPattern := `^-?\d+(\.\d+)?,-?\d+(\.\d+)?$`
	assert.Regexp(t, coordPattern, coordinates)
}

func Test_RequestAwesomeAPI_MustFailWithInvalidCEP(t *testing.T) {
	gateway := &location.AwesomeAPILocationGateway{}
	coordinates, err := gateway.Cep2Coordinates("0100100")
	assert.NotNil(t, err)
	assert.Empty(t, coordinates)
	log.Println(err)
}
