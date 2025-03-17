package usecase_test

import (
	"context"
	w "otellab/internal/gateway/weather"
	g "otellab/internal/usecase"
	"otellab/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetWeather_FailsWhenInputIsInvalid(t *testing.T) {
	input := g.WeatherByCepInputDTO{
		Cep: "010010-01",
	}

	ctx := context.Background()
	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCEP", ctx, "010010-01").Return(w.WeatherOutput{}, nil)

	usecase := g.NewGetWeatherByCEPUsecase(weatherMock)

	w, err := usecase.Execute(ctx, input)
	assert.NotNil(t, err)
	assert.Equal(t, w, g.WeatherByCepOutputDTO{})
	assert.Equal(t, err.Error(), "Must be in the format 01001001 for CEP")
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCEP", 0)
}

func Test_GetWeather_FailsWhenAPIFails(t *testing.T) {
	input := g.WeatherByCepInputDTO{
		Cep: "01001001",
	}

	ctx := context.Background()
	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCEP", ctx, "01001001").Return(w.WeatherOutput{}, assert.AnError)

	usecase := g.NewGetWeatherByCEPUsecase(weatherMock)

	w, err := usecase.Execute(ctx, input)
	assert.NotNil(t, err)
	assert.Equal(t, w, g.WeatherByCepOutputDTO{})
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCEP", 1)
}

func Test_GetWeather_Success(t *testing.T) {
	input := g.WeatherByCepInputDTO{
		Cep: "01001001",
	}
	ctx := context.Background()

	weatherMock := &mocks.MockWeatherGateway{}
	weatherMock.On("GetWeatherByCEP", ctx, "01001001").Return(w.WeatherOutput{
		City: "Sao Paulo",
	}, nil)

	usecase := g.NewGetWeatherByCEPUsecase(weatherMock)

	w, err := usecase.Execute(ctx, input)
	assert.Nil(t, err)
	assert.Equal(t, w.City, "Sao Paulo")
	weatherMock.AssertNumberOfCalls(t, "GetWeatherByCEP", 1)
}
