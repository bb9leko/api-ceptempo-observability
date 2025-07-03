package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
)

type CepRequest struct {
	Cep string `json:"cep"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func isValidCep(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func ConsultaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req CepRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if !isValidCep(req.Cep) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "invalid zipcode"})
		return
	}

	// Consulta a API original usando GET
	apiURL := os.Getenv("API_CEP_TEMPO_URL") + "?cep=" + req.Cep
	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Erro ao consultar API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
