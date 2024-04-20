package exporter

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func NewConsoleExporter() (trace.SpanExporter, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}
	return exporter, err
}
