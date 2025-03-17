package weather

import (
	weatherAPI "temperatures/pkg/weather_api"
)

type WeatherAPIGateway struct{}

func (w *WeatherAPIGateway) GetWeatherByCoordinates(coordinates string) (weather WeatherOutput, err error) {
	weatherData, err := weatherAPI.RequestWeatherAPI(coordinates)
	if err != nil {
		return weather, err
	}

	weather = WeatherOutput{
		Coordinates:     coordinates,
		Temperature:     weatherData.Current.TempC,
		TemperatureUnit: "Celsius",
		City:            weatherData.Location.Name,
		Region:          weatherData.Location.Region,
		Country:         weatherData.Location.Country,
		Pressure:        weatherData.Current.PressureMb,
		PressureUnit:    "millibars",
		Humidity:        weatherData.Current.Humidity,
		WindSpeed:       weatherData.Current.WindKph,
		WindSpeedUnit:   "kilometers per hour",
		WindDirection:   weatherData.Current.WindDir,
	}
	return weather, nil
}
