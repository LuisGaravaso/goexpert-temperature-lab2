package weather_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TimezoneID     string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	LastUpdated      string    `json:"last_updated"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	IsDay            int       `json:"is_day"`
	Condition        Condition `json:"condition"`
	WindMph          float64   `json:"wind_mph"`
	WindKph          float64   `json:"wind_kph"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	PressureMb       float64   `json:"pressure_mb"`
	PressureIn       float64   `json:"pressure_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PrecipIn         float64   `json:"precip_in"`
	Humidity         int       `json:"humidity"`
	Cloud            int       `json:"cloud"`
	FeelsLikeC       float64   `json:"feelslike_c"`
	FeelsLikeF       float64   `json:"feelslike_f"`
	WindChillC       float64   `json:"windchill_c"`
	WindChillF       float64   `json:"windchill_f"`
	HeatIndexC       float64   `json:"heatindex_c"`
	HeatIndexF       float64   `json:"heatindex_f"`
	DewPointC        float64   `json:"dewpoint_c"`
	DewPointF        float64   `json:"dewpoint_f"`
	VisibilityKm     float64   `json:"vis_km"`
	VisibilityMiles  float64   `json:"vis_miles"`
	UV               float64   `json:"uv"`
	GustMph          float64   `json:"gust_mph"`
	GustKph          float64   `json:"gust_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

func GetWeatherAPIKey() (string, error) {

	log.Printf("Loading Configs")
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey != "" {
		return apiKey, nil
	}

	configs, err := LoadConfig()
	if err != nil {
		return "", err
	}

	return configs.WeatherAPIKey, nil
}

func BuildRequestURL(api_key, coordinates string) string {
	return fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", api_key, coordinates)
}

// RequestWeatherAPI makes a request to the weather API using the provided coordinates
func RequestWeatherAPI(coordinates string) (WeatherResponse, error) {
	weatherAPIKey, err := GetWeatherAPIKey()
	if err != nil {
		return WeatherResponse{}, errors.New("error: could not get weather API key")
	}

	requestURL := BuildRequestURL(weatherAPIKey, coordinates)
	log.Println("Requesting weather API with URL:", requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, fmt.Errorf("error: received non-200 status code %d", resp.StatusCode)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return WeatherResponse{}, err
	}

	return weather, nil
}
