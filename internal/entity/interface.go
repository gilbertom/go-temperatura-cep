package entity

// CepRepository represents a repository for CEP data.
type CepRepository interface {
    GetLocalityByCep(cep string) (*Cep, error)
}

// WeatherRepository represents a repository for weather data.
type WeatherRepository interface {
    GetTemperaturesByLocality(locality string) (*Weather, error)
}