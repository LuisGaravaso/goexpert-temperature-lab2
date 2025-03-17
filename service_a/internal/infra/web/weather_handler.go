package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	weatherGateway "otellab/internal/gateway/weather"
	usecase "otellab/internal/usecase"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// Error422 representa um erro de validação de dados (unprocessable entity).
type Error422 struct {
	Error string `json:"error" example:"invalid zipcode"`
}

// Error400 representa um erro de requisição malformada (bad request).
type Error400 struct {
	Error string `json:"error" example:"invalid JSON body, must be in the format {\"cep\": \"01001001\"}"`
}

// Error404 representa um erro quando a localização não é encontrada.
type Error404 struct {
	Error string `json:"error" example:"location not found"`
}

// Error500 representa um erro interno do servidor.
type Error500 struct {
	Error string `json:"error" example:"internal server error"`
}

// WebWeatherHandler é o handler responsável por lidar com as requisições de clima.
type WebWeatherHandler struct {
	WeatherGateway weatherGateway.WeatherGateway
}

// NewWebWeatherHandler cria uma nova instância de WebWeatherHandler.
func NewWebWeatherHandler(
	weatherGateway weatherGateway.WeatherGateway,
) *WebWeatherHandler {
	return &WebWeatherHandler{
		WeatherGateway: weatherGateway,
	}
}

// Location representa a estrutura esperada no corpo da requisição.
type Location struct {
	Cep string `json:"cep" example:"01001001"`
}

// Post godoc
// @Summary Get Weather by CEP
// @Description Recebe um CEP no corpo da requisição e retorna informações de clima.
// @Tags Weather
// @Accept json
// @Produce json
// @Param request body Location true "CEP da Localização"
// @Success 200 {object} usecase.WeatherByCepOutputDTO
// @Failure 400 {object} Error400
// @Failure 404 {object} Error404
// @Failure 422 {object} Error422
// @Failure 500 {object} Error500
// @Router /temperature [post]
func (h *WebWeatherHandler) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	serviceName := os.Getenv("SERVICE_NAME")
	tr := otel.Tracer(serviceName)

	// ➕ Span: Handler Principal
	ctx, handlerSpan := tr.Start(ctx, "POST /temperature")
	defer handlerSpan.End()

	// ➕ Span: Parse do JSON
	ctx, parseSpan := tr.Start(ctx, "parse-request-body")
	parseSpan.AddEvent("Start decoding JSON body")

	var location Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		parseSpan.RecordError(err)
		parseSpan.SetStatus(codes.Error, "failed to decode JSON body")
		parseSpan.SetAttributes(attribute.String("error", err.Error()))

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error400{Error: "invalid JSON body, must be in the format {cep: 01001001}"})
		parseSpan.End()
		return
	}

	if location.Cep == "" {
		parseSpan.SetStatus(codes.Error, "missing or empty CEP")
		parseSpan.SetAttributes(attribute.String("error", "missing or empty cep"))

		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(Error400{Error: "invalid JSON body, must be in the format {cep: 01001001}"})
		parseSpan.End()
		return
	}

	parseSpan.AddEvent("Successfully decoded JSON body")
	parseSpan.SetAttributes(attribute.String("cep", location.Cep))
	parseSpan.End() // ✅ End here before starting next span

	// ➕ Span: Execução do UseCase
	ctx, usecaseSpan := tr.Start(ctx, "execute-usecase")
	usecaseSpan.AddEvent("Start executing usecase")

	dto := usecase.WeatherByCepInputDTO{Cep: location.Cep}
	useCase := usecase.NewGetWeatherByCEPUsecase(h.WeatherGateway)

	output, err := useCase.Execute(ctx, dto)
	if err != nil && err.Error() == "Must be in the format 01001001 for CEP" {
		usecaseSpan.RecordError(err)
		usecaseSpan.SetStatus(codes.Error, "invalid zipcode format")
		usecaseSpan.SetAttributes(attribute.String("error", err.Error()))

		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(Error422{Error: "invalid zipcode"})
		usecaseSpan.End()
		return
	}

	if output.Coordinates == "" {
		usecaseSpan.SetStatus(codes.Error, "location not found")
		usecaseSpan.AddEvent("Error: location not found")

		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(Error404{Error: "location not found"})
		usecaseSpan.End()
		return
	}

	usecaseSpan.AddEvent("Successfully completed usecase execution")
	usecaseSpan.SetAttributes(
		attribute.String("output.coordinates", output.Coordinates),
		attribute.String("output.temperature", fmt.Sprintf("%f", output.TemperatureCelsius)),
	)
	usecaseSpan.End() // ✅ End before response step

	// ➕ Span: Retorno da resposta
	_, responseSpan := tr.Start(ctx, "return-response")
	responseSpan.AddEvent("Start preparing response")

	w.Header().Set("Content-Type", "application/json")
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

	responseSpan.AddEvent("Successfully encoded and returned response")
	responseSpan.End()
}
