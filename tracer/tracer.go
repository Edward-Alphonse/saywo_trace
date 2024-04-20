package tracer

import (
	"context"
	"github.com/Edward-Alphonse/saywo_trace/provider"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	libraryName = "github.com/Edward-Alphonse/saywo_trace"
)

var tracer = otel.Tracer(libraryName)

type ProviderType string

const (
	ProviderTypeConsole = "console"
	ProviderTypeFile    = "file"
	ProviderTypeJaeger  = "jaeger"
)

type Config struct {
	ServerName string
	Type       ProviderType
	Endpoint   string
	File       string
}

func Init(config *Config) {
	serverName := config.ServerName
	switch config.Type {
	case ProviderTypeFile:
		provider.InitFileProvider(serverName, config.File)
	case ProviderTypeConsole:
		provider.InitConsoleProvider(serverName)
	case ProviderTypeJaeger:
		provider.InitJaegerProvider(serverName, config.Endpoint)
	}
}

func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName, opts...)
}
