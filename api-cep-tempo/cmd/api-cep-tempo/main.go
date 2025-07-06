package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bb9leko/api-cep-tempo/internal/handler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func initTracer() func() {
	exporter, err := zipkin.New(os.Getenv("ZIPKIN_URL"))
	if err != nil {
		log.Fatalf("erro ao criar exporter zipkin: %v", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	return func() { _ = tp.Shutdown(context.Background()) }
}

func main() {
	shutdown := initTracer()
	defer shutdown()

	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == "" {
		log.Fatal("WEATHERAPI_KEY não configurada")
	}
	fmt.Println("WEATHERAPI_KEY carregada:", apiKey)
	//cfg, err := configs.LoadConfig()
	//if err != nil {
	//	log.Fatalf("Erro ao carregar configuração: %v", err)
	//}
	//os.Setenv("WEATHERAPI_KEY", cfg.WeatherAPIKey)
	//fmt.Println("WEATHERAPI_KEY carregada:", cfg.WeatherAPIKey)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.Handle("/cep", otelhttp.NewHandler(http.HandlerFunc(handler.CEPHandler), "CEPHandler"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
