package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gilbertom/go-temperatura-cep/internal/config"
	"github.com/gilbertom/go-temperatura-cep/internal/entity"
)

// WeatherRepository is a repository for retrieving weather data.
type WeatherRepository struct{}

// NewWeatherRepository creates a new instance of WeatherRepository.
func NewWeatherRepository() *WeatherRepository {
    return &WeatherRepository{}
}

// GetTemperaturesByLocality retrieves the temperatures for a given locality.
func (r *WeatherRepository) GetTemperaturesByLocality(locality string) (*entity.Weather, error) {
    var weather entity.Weather
    url := fmt.Sprintf("%s?q=%s&lang=pt&key=%s", config.AppConfig.URLWeather, url.QueryEscape(locality), config.AppConfig.APIKeyWeather)
    resp, err := http.Get(url)
    if err != nil {
        return &weather, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return &weather, errors.New("failed to fetch weather data")
    }

    if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
        return &weather, err
    }

    return &weather, nil
}
