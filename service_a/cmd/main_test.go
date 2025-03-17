package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	weatherGateway "otellab/internal/gateway/weather"
	"otellab/internal/infra/web"
	"otellab/internal/infra/web/webserver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var serverAddr = "http://localhost:8080"

func TestMain(m *testing.M) {
	// Start server in background
	go func() {
		weatherGateway := &weatherGateway.WeatherAPIGateway{}

		// WebServer
		webOrderHandler := web.NewWebWeatherHandler(weatherGateway)
		webServer := webserver.NewWebServer(":8080")
		webServer.AddHandler("/temperature", webOrderHandler.Post)
		log.Println("Starting web server on port", ":8080")
		webServer.Start()
	}()

	// Wait a bit for the server to boot up
	time.Sleep(500 * time.Millisecond)

	// Run the tests
	code := m.Run()

	// Cleanup logic if needed
	os.Exit(code)
}

type Location struct {
	Cep string `json:"cep"`
}

type InvalidPayload struct {
	Email string `json:"email"`
}

func Test_PostWeather_MustReturnFromValidCep(t *testing.T) {

	loc := Location{Cep: "01001001"}

	// Convert struct to JSON
	jsonData, err := json.Marshal(loc)
	if err != nil {
		panic(err)
	}

	// Make the request
	resp, err := http.Post(serverAddr+"/temperature", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Check if the response is OK
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, body)
	assert.Contains(t, string(body), "cep")
	assert.Contains(t, string(body), "coordinates")
	assert.Contains(t, string(body), "city")
	assert.Contains(t, string(body), "region")
	assert.Contains(t, string(body), "country")
	assert.Contains(t, string(body), "temp_C")
	assert.Contains(t, string(body), "temp_F")
	assert.Contains(t, string(body), "temp_K")

}

func Test_PostWeather_MustReturnFromInvalidCep(t *testing.T) {

	loc := Location{Cep: "00000000"}

	// Convert struct to JSON
	jsonData, err := json.Marshal(loc)
	if err != nil {
		panic(err)
	}

	// Make the request
	resp, err := http.Post(serverAddr+"/temperature", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Check if the response is OK
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.NotEmpty(t, body)
	assert.Contains(t, string(body), "location not found")

}

func Test_PostWeather_MustReturnFromInvalidPayload(t *testing.T) {

	loc := InvalidPayload{Email: "x.x@gnail.com"}

	// Convert struct to JSON
	jsonData, err := json.Marshal(loc)
	if err != nil {
		panic(err)
	}

	// Make the request
	resp, err := http.Post(serverAddr+"/temperature", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Check if the response is OK
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.NotEmpty(t, body)
	assert.Contains(t, string(body), "invalid JSON body, must be in the format {cep: 01001001}")

}
