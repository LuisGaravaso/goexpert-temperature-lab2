package getweather_test

import (
	w "temperatures/internal/gateway/weather"
	g "temperatures/internal/usecase/get_weather"
	"temperatures/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetWeather_FailsWhenInputIsInvalid(t *testing.T) {
	input := g.GetWeatherInputDTO{
		Location: "010010-01",
	}

	locationMock := &mocks.MockLocationGateway{}
	locationMock.On("Cep2Coordinates", "010010-01").Return("", nil)

	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCoordinates", "").Return(w.WeatherOutput{}, nil)

	usecase := g.NewGetWeatherUseCase(locationMock, weatherMock)

	w, err := usecase.Execute(input)
	assert.NotNil(t, err)
	assert.Equal(t, w, g.GetWeatherOutputDTO{})
	assert.Equal(t, err.Error(), "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates")
	locationMock.AssertNumberOfCalls(t, "Cep2Coordinates", 0)
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCoordinates", 0)
}

func Test_GetWeather_WhenInputIsCEPConvertsToCoordinates(t *testing.T) {
	input := g.GetWeatherInputDTO{
		Location: "01001001",
	}

	locationMock := &mocks.MockLocationGateway{}
	locationMock.On("Cep2Coordinates", "01001001").Return("-23.55028,-46.63389", nil)

	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCoordinates", "-23.55028,-46.63389").Return(w.WeatherOutput{}, nil)

	usecase := g.NewGetWeatherUseCase(locationMock, weatherMock)
	w, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, w.Coordinates)
	assert.Equal(t, w.Coordinates, "-23.55028,-46.63389")
	locationMock.AssertNumberOfCalls(t, "Cep2Coordinates", 1)
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCoordinates", 1)
}

func Test_GetWeather_WhenInputIsCoordinates(t *testing.T) {
	input := g.GetWeatherInputDTO{
		Location: "-23.55028,-46.63389",
	}

	locationMock := &mocks.MockLocationGateway{}
	locationMock.On("Cep2Coordinates", "-23.55028,-46.63389").Return("", nil)

	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCoordinates", "-23.55028,-46.63389").Return(w.WeatherOutput{}, nil)

	usecase := g.NewGetWeatherUseCase(locationMock, weatherMock)
	w, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotEmpty(t, w.Coordinates)
	assert.Equal(t, w.Coordinates, "-23.55028,-46.63389")
	locationMock.AssertNumberOfCalls(t, "Cep2Coordinates", 0)
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCoordinates", 1)
}

func Test_GetWeather_WhenWeatherGatewayFails(t *testing.T) {
	input := g.GetWeatherInputDTO{
		Location: "-23.55028,-46.63389",
	}

	locationMock := &mocks.MockLocationGateway{}
	locationMock.On("Cep2Coordinates", "-23.55028,-46.63389").Return("", nil)

	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCoordinates", "-23.55028,-46.63389").Return(w.WeatherOutput{}, assert.AnError)

	usecase := g.NewGetWeatherUseCase(locationMock, weatherMock)
	w, err := usecase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, w, g.GetWeatherOutputDTO{})
	locationMock.AssertNumberOfCalls(t, "Cep2Coordinates", 0)
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCoordinates", 1)
}

func Test_CelsiusToFahrenheit(t *testing.T) {
	assert.InDelta(t, g.CelsiusToFahrenheit(0), 32.0, 0.01)
	assert.InDelta(t, g.CelsiusToFahrenheit(100), 212.0, 0.01)
	assert.InDelta(t, g.CelsiusToFahrenheit(-40), -40.0, 0.01)
}

func Test_CelsiusToKelvin(t *testing.T) {

	assert.InDelta(t, g.CelsiusToKelvin(0), 273.15, 0.01)
	assert.InDelta(t, g.CelsiusToKelvin(100), 373.15, 0.01)
	assert.InDelta(t, g.CelsiusToKelvin(-40), 233.15, 0.01)
}
