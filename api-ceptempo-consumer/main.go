package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

	http.Handle("/cep", otelhttp.NewHandler(http.HandlerFunc(ConsultaCepHandler), "ConsultaCepHandler"))
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("api-ceptempo-consumer rodando na porta 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
