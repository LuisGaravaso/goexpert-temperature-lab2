package weather_api_test

import (
	weatherapi "temperatures/pkg/weather_api"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAPIKey_MustReturnAPIKey(t *testing.T) {
	apiKey, err := weatherapi.GetWeatherAPIKey()
	assert.NoError(t, err)
	assert.NotEmpty(t, apiKey)
}

func Test_GetRequestURL_MustReturnRequestURL(t *testing.T) {
	apiKey, err := weatherapi.GetWeatherAPIKey()
	assert.NoError(t, err)

	coordinates := "40.7128,-74.0060"
	requestURL := weatherapi.BuildRequestURL(apiKey, coordinates)
	assert.NotEmpty(t, requestURL)
	assert.Contains(t, requestURL, apiKey)
	assert.Contains(t, requestURL, coordinates)
}

func Test_RequestWeatherAPI_MustReturnWeatherResponse(t *testing.T) {
	coordinates := "-23.533,-46.617"
	weather, err := weatherapi.RequestWeatherAPI(coordinates)
	assert.NoError(t, err)
	assert.Equal(t, weather.Location.Name, "Sao Paulo")
	assert.NotEmpty(t, weather.Current.TempC)
}
