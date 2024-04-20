package provider

import (
	"github.com/Edward-Alphonse/saywo_trace/exporter"
	"github.com/Edward-Alphonse/saywo_trace/resource"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
func InitJaegerProvider(serverName, endpoint string) {
	provider := NewJaegerProvider(serverName, endpoint)

	// 设置全局 TracerProvider
	otel.SetTracerProvider(provider)
}

func NewJaegerProvider(serverName, endpoint string) *trace.TracerProvider {
	// 创建 Jaeger 导出器
	exp, err := exporter.NewJaegerExporter(endpoint)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	// 创建 TracerProvider
	rsc := resource.NewResource(serverName)
	provider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(trace.NewSimpleSpanProcessor(exp)),
		//将服务实例数据添加到 tracer
		trace.WithResource(rsc),
	)
	return provider
}
