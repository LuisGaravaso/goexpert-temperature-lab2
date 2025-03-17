package weather

import "context"

type WeatherOutput struct {
	Coordinates           string
	City                  string
	Region                string
	Country               string
	TemperatureCelsius    float64
	TemperatureFahrenheit float64
	TemperatureKelvin     float64
}

type WeatherGateway interface {
	GetWeatherByCEP(ctx context.Context, coordinates string) (weather WeatherOutput, err error)
}
