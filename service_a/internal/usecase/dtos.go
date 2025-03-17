package usecase

type WeatherByCepInputDTO struct {
	Cep string
}

type WeatherByCepOutputDTO struct {
	Cep                   string  `json:"cep"`
	Coordinates           string  `json:"coordinates"`
	City                  string  `json:"city"`
	Region                string  `json:"region"`
	Country               string  `json:"country"`
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}
