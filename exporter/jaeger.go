package exporter

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func NewJaegerExporter(endpoint string) (trace.SpanExporter, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}
	return exporter, err
}
