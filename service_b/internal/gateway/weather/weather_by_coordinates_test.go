package weather_test

import (
	"log"
	"testing"

	"temperatures/internal/gateway/weather"

	"github.com/stretchr/testify/assert"
)

func Test_RequestWeatherAPI_MustReturnWeatherData(t *testing.T) {
	service := &weather.WeatherAPIGateway{}
	weatherData, err := service.GetWeatherByCoordinates("-23.5489,-46.6388")
	assert.Nil(t, err)
	assert.NotEmpty(t, weatherData)
	log.Println(weatherData)
}
