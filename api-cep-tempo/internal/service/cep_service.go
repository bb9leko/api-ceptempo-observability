package service

import (
	"math"
	"regexp"

	"github.com/bb9leko/api-cep-tempo/internal/client"
	"github.com/bb9leko/api-cep-tempo/internal/model"
)

// ServiceError representa um erro customizado com código HTTP e mensagem
type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func IsValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func GetCEPAndTempoInfo(cep string) (*model.CEPTempoResponse, error) {
	if !IsValidCEP(cep) {
		return nil, &ServiceError{Code: 422, Message: "invalid zipcode"}
	}
	data, err := client.FetchCEP(cep)
	if err != nil {
		return nil, &ServiceError{Code: 500, Message: "internal error"}
	}
	if data.Erro {
		return nil, &ServiceError{Code: 404, Message: "can not find zipcode"}
	}

	// Consulta a WeatherAPI usando a localidade retornada
	tempC, err := client.FetchTemperatura(data.Localidade)
	if err != nil {
		return nil, err
	}

	tempF := tempC*1.8 + 32 // Fahrenheit
	tempK := tempC + 273    // Kelvin (conforme solicitado, sem casas decimais)

	// Arredonda para 1 casa decimal
	tempC = math.Round(tempC*10) / 10
	tempF = math.Round(tempF*10) / 10
	tempK = math.Round(tempK*10) / 10

	return &model.CEPTempoResponse{
		Cep:        data.Cep,
		Localidade: data.Localidade,
		Temperatura: model.TempoResponse{
			Celsius:    tempC,
			Fahrenheit: tempF,
			Kelvin:     tempK,
		},
	}, nil
}
