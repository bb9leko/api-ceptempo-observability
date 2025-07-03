package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bb9leko/api-cep-tempo/internal/model"
)

func FetchCEP(cep string) (*model.ViaCEPResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data model.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
