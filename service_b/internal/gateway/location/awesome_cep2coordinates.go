package location

import (
	"errors"
	awesome "temperatures/pkg/awesome_api"
)

type AwesomeAPILocationGateway struct{}

func (a *AwesomeAPILocationGateway) Cep2Coordinates(cep string) (coordinates string, err error) {
	location := awesome.RequestAwesomeAPI(cep)
	if location.Latitude == "" || location.Longitude == "" {
		return "", errors.New("invalid CEP")
	}
	coordinates = location.Latitude + "," + location.Longitude
	return coordinates, nil
}
