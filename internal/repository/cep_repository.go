package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gilbertom/go-temperatura-cep/internal/config"
	"github.com/gilbertom/go-temperatura-cep/internal/entity"
)

// CepRepository represents a repository for handling CEP data.
type CepRepository struct{}

// NewCepRepository creates a new instance of CepRepository.
func NewCepRepository() *CepRepository {
    return &CepRepository{}
}

// GetLocalityByCep retrieves the locality information for a given CEP.
func (r *CepRepository) GetLocalityByCep(cep string) (*entity.Cep, error) {
    var locality entity.Cep
    url := fmt.Sprintf("%s/%s/json/", config.AppConfig.URLCep, cep)
    resp, err := http.Get(url)
    if err != nil {
        return &locality, err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusBadRequest {
        return &locality, errors.New("invalid zipcode")
    }

    if err := json.NewDecoder(resp.Body).Decode(&locality); err != nil {
        return &locality, err
    }

    if locality.Erro == "true" {
        return &locality, errors.New("can not find zipcode")
    }

    if locality.Localidade == "" {
        return &locality, errors.New("invalid zipcode")
    }

    return &locality, nil
}
