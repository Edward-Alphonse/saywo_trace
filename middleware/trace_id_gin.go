package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
)

func WithTraceID(c *gin.Context) {
	span := trace.SpanFromContext(c.Request.Context())
	traceID := span.SpanContext().TraceID().String()
	c.Header("x-trace-id", traceID)
	c.Next()
}

func WithOtelTrace(hostName string) gin.HandlerFunc {
	return otelgin.Middleware(hostName)
}
