package weather_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Weather struct {
	Coordinates                string  `json:"coordinates"`
	City                       string  `json:"city"`
	Region                     string  `json:"region"`
	Country                    string  `json:"country"`
	TempC                      float64 `json:"temp_C"`
	TempF                      float64 `json:"temp_F"`
	TempK                      float64 `json:"temp_K"`
	PressureInMillibars        int     `json:"pressure_in_millibars"`
	PrecipitationInMillimeters int     `json:"precipitation_in_millimeters"`
	HumidityInPercentage       int     `json:"humidity_in_percentage"`
	WindInKph                  float64 `json:"wind_in_kph"`
	WindDirection              string  `json:"wind_direction"`
}

func RequestWeatherAPI(ctx context.Context, cep string) Weather {
	weatherUrl := os.Getenv("WEATHER_API_URL")
	if weatherUrl == "" {
		weatherUrl = "http://localhost:8081/temperature"
	}
	url := weatherUrl + "/" + cep

	// Create a new instrumented HTTP client
	client := http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Create request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode the response
	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatal(err)
	}

	return weather
}
