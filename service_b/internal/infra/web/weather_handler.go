package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	locationGateway "temperatures/internal/gateway/location"
	weatherGateway "temperatures/internal/gateway/weather"
	usecase "temperatures/internal/usecase/get_weather"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type Error422 struct {
	Error string `json:"error" example:"invalid location"`
}

type Error404 struct {
	Error string `json:"error" example:"location not found"`
}

type Error500 struct {
	Error string `json:"error" example:"internal server error"`
}

type WebWeatherHandler struct {
	LocationGateway locationGateway.LocationGateway
	WeatherGateway  weatherGateway.WeatherGateway
}

func NewWebWeatherHandler(
	locationGateway locationGateway.LocationGateway,
	weatherGateway weatherGateway.WeatherGateway,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		LocationGateway: locationGateway,
		WeatherGateway:  weatherGateway,
	}
}

// Get godoc
// @Summary Get Weather by Location
// @Description Returns weather info by location (CEP or lat,lng)
// @Tags Weather
// @Accept json
// @Produce json
// @Param location path string true "Location (CEP or lat,lng)"
// @Success 200 {object} usecase.GetWeatherOutputDTO
// @Failure 404 {object} Error404
// @Failure 422 {object} Error422
// @Failure 500 {object} Error500
// @Router /temperature/{location} [get]
func (h *WebWeatherHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	serviceName := os.Getenv("SERVICE_NAME")
	tr := otel.Tracer(serviceName)

	// ➕ Span: Handler Principal
	ctx, handlerSpan := tr.Start(ctx, "GET /temperature/{location}")
	defer handlerSpan.End()

	w.Header().Set("Content-Type", "application/json")

	// ➕ Span: Extração da localização da URL
	ctx, extractSpan := tr.Start(ctx, "extract-location-from-url")
	extractSpan.AddEvent("Extract location from URL")
	location := chi.URLParam(r, "location")
	extractSpan.SetAttributes(attribute.String("input.location", location))
	extractSpan.End()

	dto := usecase.GetWeatherInputDTO{Location: location}

	// ➕ Span: Execução do UseCase
	ctx, usecaseSpan := tr.Start(ctx, "execute-usecase")
	usecaseSpan.AddEvent("Start executing weather usecase")
	useCase := usecase.NewGetWeatherUseCase(h.LocationGateway, h.WeatherGateway)
	output, err := useCase.Execute(dto)

	if err != nil && err.Error() == "Must be in the format 01001001 for CEP or -23.55028,-46.63389 for Coordinates" {
		usecaseSpan.RecordError(err)
		usecaseSpan.SetStatus(codes.Error, "invalid location format")
		usecaseSpan.SetAttributes(attribute.String("error", err.Error()))
		usecaseSpan.End()

		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(Error422{Error: "invalid location"})
		return
	}

	if output.Coordinates == "" {
		usecaseSpan.SetStatus(codes.Error, "location not found")
		usecaseSpan.AddEvent("Location not found")
		usecaseSpan.End()

		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(Error404{Error: "location not found"})
		return
	}

	usecaseSpan.AddEvent("Successfully completed usecase execution")
	usecaseSpan.SetAttributes(
		attribute.String("output.coordinates", output.Coordinates),
		attribute.String("output.temperature", fmt.Sprintf("%f", output.TemperatureInCelsius)),
	)
	usecaseSpan.End()

	// ➕ Span: Preparando resposta
	_, responseSpan := tr.Start(ctx, "return-response")
	responseSpan.AddEvent("Start encoding response")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		responseSpan.RecordError(err)
		responseSpan.SetStatus(codes.Error, "failed to encode response")
		responseSpan.SetAttributes(attribute.String("error", err.Error()))

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(Error500{Error: "internal server error"})
		responseSpan.End()
		return
	}

	responseSpan.AddEvent("Successfully encoded response")
	responseSpan.End()
}
