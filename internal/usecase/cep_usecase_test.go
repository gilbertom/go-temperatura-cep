package usecase

import (
	"errors"
	"testing"

	"github.com/gilbertom/go-temperatura-cep/internal/entity"
)

// Mock Cep Repository
type MockCepRepository struct{}

func (m *MockCepRepository) GetLocalityByCep(cep string) (*entity.Cep, error) {
	if cep == "88888888" || len(cep) != 8 || len(cep) == 8 && !isNumeric(cep) {
		return &entity.Cep{Erro: "invalid CEP"}, errors.New("invalid CEP")
	}
	localities := map[string]string{
		"01001000": "São Paulo",
		"28951620": "Cabo Frio",
	}
	if locality, exists := localities[cep]; exists {
		return &entity.Cep{Localidade: locality, Erro: ""}, nil
	}
	return &entity.Cep{Erro: "invalid CEP"}, errors.New("invalid CEP")
}

func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func TestGetLocalityByCep(t *testing.T) {
	mockRepo := &MockCepRepository{}
	cepUsecase := NewCepUsecase(mockRepo)

	tests := []struct {
		input    string
		expected string
		err      string
	}{
		{"01001000", "São Paulo", ""},
		{"28951620", "Cabo Frio", ""},
		{"88888888", "", "invalid CEP"},
		{"12345678", "", "invalid CEP"},
		{"1234567A", "", "invalid CEP"},
		{"1234567", "", "invalid CEP"},
		{"123456789", "", "invalid CEP"},
	}

	for _, test := range tests {
		locality, err := cepUsecase.GetLocalityByCep(test.input)
		if err != nil && err.Error() != test.err {
			t.Fatalf("for CEP %v, expected error %v, got %v", test.input, test.err, err)
		}
		if locality.Localidade != test.expected {
			t.Fatalf("for CEP %v, expected locality %v, got %v", test.input, test.expected, locality.Localidade)
		}
		if locality.Erro != test.err {
			t.Fatalf("for CEP %v, expected error message %v, got %v", test.input, test.err, locality.Erro)
		}
	}
}
