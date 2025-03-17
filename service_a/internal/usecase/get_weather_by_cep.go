package usecase

import (
	"context"
	"errors"
	"otellab/internal/entities"
	"otellab/internal/gateway/weather"
)

type GetWeatherByCEPUsecase struct {
	WeatherGateway weather.WeatherGateway
}

func NewGetWeatherByCEPUsecase(weatherGateway weather.WeatherGateway) GetWeatherByCEPUsecase {
	return GetWeatherByCEPUsecase{WeatherGateway: weatherGateway}
}

func (u *GetWeatherByCEPUsecase) Execute(ctx context.Context, input WeatherByCepInputDTO) (WeatherByCepOutputDTO, error) {

	location := entities.NewLocation(input.Cep)
	if !location.IsValid {
		return WeatherByCepOutputDTO{}, errors.New(location.InvalidMessage)
	}

	weatherOutput, err := u.WeatherGateway.GetWeatherByCEP(ctx, location.Id)
	if err != nil {
		return WeatherByCepOutputDTO{}, err
	}

	return WeatherByCepOutputDTO{
		Cep:                   input.Cep,
		Coordinates:           weatherOutput.Coordinates,
		City:                  weatherOutput.City,
		Region:                weatherOutput.Region,
		Country:               weatherOutput.Country,
		TemperatureCelsius:    weatherOutput.TemperatureCelsius,
		TemperatureFahrenheit: weatherOutput.TemperatureFahrenheit,
		TemperatureKelvin:     weatherOutput.TemperatureKelvin,
	}, nil
}
