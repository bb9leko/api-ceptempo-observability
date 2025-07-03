package model

type ViaCEPResponse struct {
	Cep        string `json:"cep"`
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro,omitempty"`
}
