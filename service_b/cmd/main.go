// @title Weather API
// @version 1.0
// @description API for retrieving weather information by location
// @BasePath /

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "temperatures/docs"
	locationGateway "temperatures/internal/gateway/location"
	weatherGateway "temperatures/internal/gateway/weather"
	"temperatures/internal/infra/observability/otel"
	"temperatures/internal/infra/web"
	"temperatures/internal/infra/web/webserver"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	locationGateway := &locationGateway.AwesomeAPILocationGateway{}
	weatherGateway := &weatherGateway.WeatherAPIGateway{}

	// Environment
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if otlpEndpoint != "" {
		// Configuração do OpenTelemetry
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt)

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		// Inicializa o OpenTelemetry
		shutdown, err := otel.InitProvider("service-b-get-weather-by-location", otlpEndpoint)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err := shutdown(ctx); err != nil {
				log.Fatalf("failed to shutdown TracerProvider: %v", err)
			}
		}()
	}

	// Porta dinâmica para compatibilidade com Cloud Run
	port := os.Getenv("WEBPORT")
	if port == "" {
		port = "8081"
	}

	// WebServer
	webWeatherHandler := web.NewWebWeatherHandler(locationGateway, weatherGateway)
	webServer := webserver.NewWebServer(":" + port)

	webServer.AddHandler(
		"/temperature/{location}",
		http.HandlerFunc(otelhttp.NewHandler(http.HandlerFunc(webWeatherHandler.Get),
			"ServiceB-PostTemperatureHandler").ServeHTTP))

	log.Println("Starting web server on port", ":"+port)
	webServer.Start()
}
