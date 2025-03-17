package weather_test

import (
	"context"
	"testing"

	"otellab/internal/gateway/weather"

	"github.com/stretchr/testify/assert"
)

func Test_RequestWeatherAPI_MustReturnWeatherData(t *testing.T) {
	service := &weather.WeatherAPIGateway{}
	ctx := context.Background()
	weatherData, err := service.GetWeatherByCEP(ctx, "01001001")
	assert.Nil(t, err)
	assert.NotEmpty(t, weatherData)
	assert.Equal(t, weatherData.City, "Sao Paulo")
}
