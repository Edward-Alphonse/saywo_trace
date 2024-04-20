package provider

import (
	"github.com/Edward-Alphonse/saywo_trace/exporter"
	"github.com/Edward-Alphonse/saywo_trace/resource"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func InitConsoleProvider(serverName string) {
	exp, err := exporter.NewConsoleExporter()
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	rsc := resource.NewResource(serverName)
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(trace.NewSimpleSpanProcessor(exp)),
		//将服务实例数据添加到 tracer
		trace.WithResource(rsc),
	)

	//注册为全局追踪程序
	otel.SetTracerProvider(tp)
}
