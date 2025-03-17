package getweather

import (
	"errors"
	"temperatures/internal/entities"
	locationGateway "temperatures/internal/gateway/location"
	weatherGateway "temperatures/internal/gateway/weather"
)

type GetWeatherUseCase struct {
	LocationGateway locationGateway.LocationGateway
	WeatherGateway  weatherGateway.WeatherGateway
}

func NewGetWeatherUseCase(
	locationGateway locationGateway.LocationGateway,
	weatherGateway weatherGateway.WeatherGateway,
) GetWeatherUseCase {
	return GetWeatherUseCase{
		LocationGateway: locationGateway,
		WeatherGateway:  weatherGateway,
	}
}

func (g *GetWeatherUseCase) Execute(inputDTO GetWeatherInputDTO) (GetWeatherOutputDTO, error) {

	//Check if the location is valid
	location := entities.NewLocation(inputDTO.Location)
	if !location.IsValid {
		return GetWeatherOutputDTO{}, errors.New(location.InvalidMessage)
	}

	//Make sure the location is in the format of coordinates
	var coordinates string
	if location.Type == "CEP" {
		coordinates, _ = g.LocationGateway.Cep2Coordinates(location.Id)
	} else {
		coordinates = location.Id
	}

	// Get the weather for these coordinates
	weather, err := g.WeatherGateway.GetWeatherByCoordinates(coordinates)
	if err != nil {
		return GetWeatherOutputDTO{}, err
	}

	//Return the current weather condition for these coordinates
	return GetWeatherOutputDTO{
		Coordinates:             coordinates,
		City:                    weather.City,
		Region:                  weather.Region,
		Country:                 weather.Country,
		TemperatureInCelsius:    weather.Temperature,
		TemperatureInFahrenheit: CelsiusToFahrenheit(weather.Temperature),
		TemperatureInKelvin:     CelsiusToKelvin(weather.Temperature),
		PressureInMillibars:     weather.Pressure,
		HumidityInPercentage:    weather.Humidity,
		WindInKph:               weather.WindSpeed,
		WindDirection:           weather.WindDirection,
	}, nil
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}
