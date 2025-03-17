package getweather

type GetWeatherInputDTO struct {
	Location string `json:"location"`
}

type GetWeatherOutputDTO struct {
	Coordinates                string  `json:"coordinates"`
	City                       string  `json:"city"`
	Region                     string  `json:"region"`
	Country                    string  `json:"country"`
	TemperatureInCelsius       float64 `json:"temp_C"`
	TemperatureInFahrenheit    float64 `json:"temp_F"`
	TemperatureInKelvin        float64 `json:"temp_K"`
	PressureInMillibars        float64 `json:"pressure_in_millibars"`
	PrecipitationInMillimeters float64 `json:"precipitation_in_millimeters"`
	HumidityInPercentage       int     `json:"humidity_in_percentage"`
	WindInKph                  float64 `json:"wind_in_kph"`
	WindDirection              string  `json:"wind_direction"`
}
