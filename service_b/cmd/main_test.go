package main

import (
	"log"
	"net/http"
	"os"
	locationGateway "temperatures/internal/gateway/location"
	weatherGateway "temperatures/internal/gateway/weather"
	"temperatures/internal/infra/web"
	"temperatures/internal/infra/web/webserver"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var serverAddr = "http://localhost:8080"

func TestMain(m *testing.M) {
	// Start server in background
	go func() {
		locationGateway := &locationGateway.AwesomeAPILocationGateway{}
		weatherGateway := &weatherGateway.WeatherAPIGateway{}

		// WebServer
		webOrderHandler := web.NewWebWeatherHandler(locationGateway, weatherGateway)
		webServer := webserver.NewWebServer(":8080")
		webServer.AddHandler("/temperature/{location}", webOrderHandler.Get)
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

func Test_GetWeather_MustReturn422ForInvalidInput(t *testing.T) {

	resp, err := http.Get(serverAddr + "/temperature/010010-01")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func Test_GetWeather_MustReturn404ForInvalidCEP(t *testing.T) {

	resp, err := http.Get(serverAddr + "/temperature/00000000")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func Test_GetWeather_MustReturn404ForInvalidCoordinates(t *testing.T) {

	resp, err := http.Get(serverAddr + "/temperature/-90,-180")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func Test_GetWeather_MustReturn200ForValidCEP(t *testing.T) {

	resp, err := http.Get(serverAddr + "/temperature/01001001")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func Test_GetWeather_MustReturn200ForValidCoordinates(t *testing.T) {

	resp, err := http.Get(serverAddr + "/temperature/-23.55028,-46.63389")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}
