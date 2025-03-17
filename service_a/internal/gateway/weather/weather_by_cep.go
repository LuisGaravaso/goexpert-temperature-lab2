package weather

import (
	"context"
	"errors"
	weatherAPI "otellab/pkg/weather_api"
)

type WeatherAPIGateway struct{}

func (w *WeatherAPIGateway) GetWeatherByCEP(ctx context.Context, cep string) (weather WeatherOutput, err error) {
	weatherData := weatherAPI.RequestWeatherAPI(ctx, cep)
	if weatherData.City == "" {
		return WeatherOutput{}, errors.New("invalid CEP")
	}

	weather = WeatherOutput{
		Coordinates:           weatherData.Coordinates,
		City:                  weatherData.City,
		Region:                weatherData.Region,
		Country:               weatherData.Country,
		TemperatureCelsius:    weatherData.TempC,
		TemperatureFahrenheit: weatherData.TempF,
		TemperatureKelvin:     weatherData.TempK,
	}
	return weather, nil
}
