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

	_ "otellab/docs"
	weatherGateway "otellab/internal/gateway/weather"
	"otellab/internal/infra/web"
	"otellab/internal/infra/web/webserver"

	"otellab/internal/infra/observability/otel"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	weatherGateway := &weatherGateway.WeatherAPIGateway{}

	// Environment
	otlpEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	serviceName := os.Getenv("SERVICE_NAME")

	// Configuração do OpenTelemetry
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Inicializa o OpenTelemetry
	shutdown, err := otel.InitProvider(serviceName, otlpEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %v", err)
		}
	}()

	// Porta dinâmica para compatibilidade com Cloud Run
	port := os.Getenv("WEBPORT")
	if port == "" {
		port = "8080"
	}

	// WebServer
	webWeatherHandler := web.NewWebWeatherHandler(weatherGateway)
	webServer := webserver.NewWebServer(":" + port)

	// Adiciona o handler com a instrumentação do OpenTelemetry
	webServer.AddHandler("/temperature",
		http.HandlerFunc(otelhttp.NewHandler(http.HandlerFunc(webWeatherHandler.Post),
			"ServiceA-PostTemperatureHandler").ServeHTTP))

	log.Println("Starting web server on port", ":"+port)
	webServer.Start()
}
