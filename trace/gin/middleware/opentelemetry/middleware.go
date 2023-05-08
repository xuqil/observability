package opentelemetry

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationName = "github.com/xuqil/observability/trace/gin/middlewares/opentelemetry"

type MiddlewareBuilder struct {
	Tracer trace.Tracer
}

func (m MiddlewareBuilder) Build() gin.HandlerFunc {
	if m.Tracer == nil {
		m.Tracer = otel.GetTracerProvider().Tracer(instrumentationName)
	}
	return func(ctx *gin.Context) {
		reqCtx := ctx.Request.Context()
		// 尝试和客户端的 trace 结合在一起
		reqCtx = otel.GetTextMapPropagator().Extract(reqCtx, propagation.HeaderCarrier(ctx.Request.Header))

		reqCtx, span := m.Tracer.Start(reqCtx, "unknown")
		defer span.End()

		span.SetAttributes(attribute.String("http.method", ctx.Request.Method))
		span.SetAttributes(attribute.String("http.url", ctx.Request.URL.String()))
		span.SetAttributes(attribute.String("http.scheme", ctx.Request.URL.Scheme))
		span.SetAttributes(attribute.String("http.host", ctx.Request.Host))

		ctx.Request = ctx.Request.WithContext(reqCtx)

		ctx.Next()

		span.SetName(ctx.FullPath())
		span.SetAttributes(attribute.Int("http.status", ctx.Writer.Status()))
	}
}
