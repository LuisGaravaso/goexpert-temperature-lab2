package weather_api_test

import (
	"context"
	"otellab/pkg/weather_api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestWeatherAPI(t *testing.T) {

	// Test with a valid CEP
	ctx := context.Background()
	weather := weather_api.RequestWeatherAPI(ctx, "01001001")
	assert.Equal(t, weather.City, "Sao Paulo")

	// Test with an invalid CEP
	ctx = context.Background()
	weather = weather_api.RequestWeatherAPI(ctx, "00000000")
	assert.Equal(t, weather.City, "")
}
