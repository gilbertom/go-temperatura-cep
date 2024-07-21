package entity

// Cep represents a CEP entity.
type Cep struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro"`
}
