package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bb9leko/api-cep-tempo/internal/service"
)

type tempResponse struct {
	Cep        string  `json:"cep,omitempty"`
	Localidade string  `json:"localidade,omitempty"`
	TempC      float64 `json:"temp_C,omitempty"`
	TempF      float64 `json:"temp_F,omitempty"`
	TempK      float64 `json:"temp_K,omitempty"`
	Code       int     `json:"code"`
	Error      string  `json:"error"`
}

func CEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	data, err := service.GetCEPAndTempoInfo(cep)
	if err != nil {
		code := 500
		msg := "internal error"
		if serr, ok := err.(*service.ServiceError); ok {
			code = serr.Code
			msg = serr.Message
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(tempResponse{
			Cep:        cep,
			Localidade: "",
			Code:       code,
			Error:      msg,
		})
		return
	}

	resp := tempResponse{
		Cep:        data.Cep,
		Localidade: data.Localidade,
		TempC:      data.Temperatura.Celsius,
		TempF:      data.Temperatura.Fahrenheit,
		TempK:      data.Temperatura.Kelvin,
		Code:       http.StatusOK,
		Error:      "",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
