package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/gilbertom/go-temperatura-cep/internal/usecase"
	"github.com/gilbertom/go-temperatura-cep/internal/web/webserver/dto"
)

// HTTPHandler handles HTTP requests.
type HTTPHandler struct {
    cepUsecase *usecase.CepUsecase
    weatherUsecase *usecase.WeatherUsecase
}

// NewHTTPHandler creates a new HTTPHandler instance.
func NewHTTPHandler(u *usecase.CepUsecase, w *usecase.WeatherUsecase) *HTTPHandler {
    return &HTTPHandler{
        cepUsecase: u,
        weatherUsecase: w,
    }
}

// GetTemperaturesByCep handles the request to get the temperatures by CEP.
func (h *HTTPHandler) GetTemperaturesByCep(w http.ResponseWriter, r *http.Request) {
    cep := r.URL.Query().Get("cep")
    
    if validCep := h.cepUsecase.ValidateCep(cep); !validCep {
        http.Error(w, "CEP is invalid", http.StatusBadRequest)
        return
    }

    locality, err := h.cepUsecase.GetLocalityByCep(cep)
    if err != nil {
        if err.Error() == "invalid zipcode" {
            http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
            return
        }

        if err.Error() == "can not find zipcode" {
            http.Error(w, "can not find zipcode", http.StatusNotFound)
            return
        }
    }

    weather, err := h.weatherUsecase.GetTemperaturesByLocality(locality.Localidade)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := dto.WeatherResponse{
        Celsius:    weather.Current.TempC,
        Fahrenheit: h.weatherUsecase.ConvertCelsiusToFahrenheit(weather.Current.TempC),
        Kelvin:     h.weatherUsecase.ConvertCelsiusToKelvin(weather.Current.TempC),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
