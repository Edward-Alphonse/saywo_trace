package provider

import (
	"github.com/Edward-Alphonse/saywo_trace/exporter"
	"github.com/Edward-Alphonse/saywo_trace/resource"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

const (
	file = "traces.txt"
)

func InitFileProvider(serverName string, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}
	//传给导出器
	exp, _ := exporter.NewFileExporter(f)
	rsc := resource.NewResource(serverName)
	tp := trace.NewTracerProvider(
		//将span添加到tracer
		trace.WithBatcher(exp),
		//将服务实例数据添加到 tracer
		trace.WithResource(rsc),
	)
	//注册为全局追踪程序
	otel.SetTracerProvider(tp)
}
