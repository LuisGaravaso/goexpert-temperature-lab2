package weather

type WeatherOutput struct {
	Coordinates     string
	Temperature     float64
	TemperatureUnit string
	City            string
	Region          string
	Country         string
	Pressure        float64
	PressureUnit    string
	Humidity        int
	WindSpeed       float64
	WindSpeedUnit   string
	WindDirection   string
}

type WeatherGateway interface {
	GetWeatherByCoordinates(coordinates string) (weather WeatherOutput, err error)
}
