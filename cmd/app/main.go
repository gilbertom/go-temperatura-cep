package main

import (
	"fmt"
	"net/http"

	"github.com/gilbertom/go-temperatura-cep/internal/config"
	"github.com/gilbertom/go-temperatura-cep/internal/repository"
	"github.com/gilbertom/go-temperatura-cep/internal/usecase"
	"github.com/gilbertom/go-temperatura-cep/internal/web/webserver"
)

func main() {
    config.LoadConfig()

    cepRepo := repository.NewCepRepository()
    weatherRepo := repository.NewWeatherRepository()

    cepUsecase := usecase.NewCepUsecase(cepRepo)
    weatherUsecase := usecase.NewWeatherUsecase(weatherRepo)

    HTTPHandler := webserver.NewHTTPHandler(cepUsecase, weatherUsecase)

    http.HandleFunc("/", HTTPHandler.GetTemperaturesByCep)
    fmt.Println("Server running on port 8080")
    if err := http.ListenAndServe(":"+config.AppConfig.PortHTTP, nil); err != nil {
        panic(err)
    }
}
