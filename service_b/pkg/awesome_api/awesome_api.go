package awesome_api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Location struct {
	Cep         string `json:"cep"`
	AddressType string `json:"address_type"`
	AddressName string `json:"address_name"`
	Address     string `json:"address"`
	State       string `json:"state"`
	District    string `json:"district"`
	Latitude    string `json:"lat"`
	Longitude   string `json:"lng"`
	City        string `json:"city"`
	CityIbge    string `json:"city_ibge"`
	Ddd         string `json:"ddd"`
}

func RequestAwesomeAPI(cep string) Location {

	// Make a GET request to the Awesome API
	resp, err := http.Get("https://cep.awesomeapi.com.br/json/" + cep)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse the response body to a Location struct
	var location Location
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		log.Fatal(err)
	}

	return location
}
