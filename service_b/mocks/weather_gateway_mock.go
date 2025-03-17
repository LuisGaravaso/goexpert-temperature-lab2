package mocks

import (
	w "temperatures/internal/gateway/weather"

	"github.com/stretchr/testify/mock"
)

type MockWeatherGateway struct {
	mock.Mock
}

func (m *MockWeatherGateway) GetWeatherByCoordinates(coordinates string) (weather w.WeatherOutput, err error) {
	args := m.Called(coordinates)
	return args.Get(0).(w.WeatherOutput), args.Error(1)
}
