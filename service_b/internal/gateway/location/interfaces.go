package location

type LocationGateway interface {
	Cep2Coordinates(cep string) (coordinates string, err error)
}
