package usecase

import "github.com/gilbertom/go-temperatura-cep/internal/entity"

// WeatherUsecase represents a use case for weather operations.
type WeatherUsecase struct {
    repo entity.WeatherRepository
}

// NewWeatherUsecase creates a new instance of WeatherUsecase.
func NewWeatherUsecase(repo entity.WeatherRepository) *WeatherUsecase {
    return &WeatherUsecase{repo: repo}
}

// GetTemperaturesByLocality retrieves the temperatures for a given locality.
func (u *WeatherUsecase) GetTemperaturesByLocality(locality string) (*entity.Weather, error) {
    return u.repo.GetTemperaturesByLocality(locality)
}

// ConvertCelsiusToFahrenheit converts a temperature value from Celsius to Fahrenheit.
func (u *WeatherUsecase) ConvertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius * 1.8 + 32
}

// ConvertCelsiusToKelvin converts a temperature value from Celsius to Kelvin.
func (u *WeatherUsecase) ConvertCelsiusToKelvin(celsius float64) float64 {
    return celsius + 273
}
