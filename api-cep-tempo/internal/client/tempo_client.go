package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

type weatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

// Remove acentos, transforma em maiúscula e substitui espaços por '+'
func sanitizeCity(city string) string {
	// Normaliza para decompor caracteres acentuados
	normCity := norm.NFD.String(city)
	t := make([]rune, 0, len(normCity))
	for _, r := range normCity {
		if unicode.Is(unicode.Mn, r) {
			continue // Remove marcas de acento
		}
		t = append(t, unicode.ToUpper(r))
	}
	return strings.ReplaceAll(string(t), " ", "+")
}

func FetchTemperatura(city string) (float64, error) {
	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("WEATHERAPI_KEY não configurada")
	}

	fmt.Println("Consultando WeatherAPI para cidade:", city) // Log do valor de city
	city = sanitizeCity(city)
	fmt.Println("Cidade sanitizada:", city) // Log após sanitização

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("WeatherAPI erro: %s - %s", resp.Status, string(body))
	}

	var data weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("Erro ao decodificar resposta WeatherAPI: %w", err)
	}
	return data.Current.TempC, nil
}
