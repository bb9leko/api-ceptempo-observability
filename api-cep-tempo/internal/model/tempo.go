package model

type TempoResponse struct {
	Celsius    float64 `json:"celsius"`
	Fahrenheit float64 `json:"fahrenheit"`
	Kelvin     float64 `json:"kelvin"`
}

type CEPTempoResponse struct {
	Cep         string        `json:"cep"`
	Localidade  string        `json:"localidade"`
	Temperatura TempoResponse `json:"temperatura"`
}
