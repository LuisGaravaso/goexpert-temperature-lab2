package mocks

import (
	"context"
	w "otellab/internal/gateway/weather"

	"github.com/stretchr/testify/mock"
)

type MockWeatherGateway struct {
	mock.Mock
}

func (m *MockWeatherGateway) GetWeatherByCEP(ctx context.Context, cep string) (weather w.WeatherOutput, err error) {
	args := m.Called(ctx, cep)
	return args.Get(0).(w.WeatherOutput), args.Error(1)
}
