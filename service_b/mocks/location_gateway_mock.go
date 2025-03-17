package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockLocationGateway struct {
	mock.Mock
}

func (m *MockLocationGateway) Cep2Coordinates(cep string) (string, error) {
	args := m.Called(cep)
	return args.String(0), args.Error(1)
}
