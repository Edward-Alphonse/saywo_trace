package main

import (
	"context"
	"github.com/Edward-Alphonse/saywo_trace/middleware"
	"github.com/Edward-Alphonse/saywo_trace/tracer"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	serverName = "my-server"
	hostName   = "my-serveraaaaaaaaa"
)

//var tracer = otel.Tracer(libraryName)

func main() {
	config := &tracer.Config{
		Type:     tracer.ProviderTypeJaeger,
		Endpoint: "http://localhost:14268/api/traces",
	}
	tracer.Init(config)

	r := gin.Default()
	r.Use(otelgin.Middleware(hostName), middleware.TraceID)
	//provider.InitConsoleTracer()

	r.GET("/hzc", func(g *gin.Context) {
		// 不使用g.request.Context() 会丢失parent span信息重新生成traceID
		newCtx, span := tracer.Start(g.Request.Context(), "Create")
		defer span.End()
		//time.Sleep(5 * time.Second)
		GetOne(newCtx, "1234")
		g.JSON(200, `{"data": 200}`)
	})
	r.Run("0.0.0.0:8080")
}

func GetOne(ctx context.Context, nStr string) {
	nCtx, span := tracer.Start(ctx, "GetOne")
	defer span.End()
	SetAge(nCtx, "hello world")
	span.SetAttributes(attribute.String("request.n", nStr))
}
func SetAge(ctx context.Context, val string) {
	_, span := tracer.Start(ctx, "SetAge", trace.WithAttributes(
		attribute.String("key3", "value3"),
		attribute.Int("key4", 456),
	))
	defer span.End()
	span.SetAttributes(attribute.String("ageTime", val))
}
