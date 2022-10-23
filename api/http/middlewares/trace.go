// Package middlewares contains all middleware for GIN
package middlewares

import (
	"auth-plus-notification/config"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// Trace is a middleware that gather trace of system
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tp, err := initTracer()
		if err != nil {
			log.Fatal(err)
		}
		otel.SetTracerProvider(tp)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// Cleanly shutdown and flush telemetry when the application exits.
		defer func(ctx context.Context) {
			// Do not make the application hang when it is shutdown.
			ctx, cancel = context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			if err := tp.Shutdown(ctx); err != nil {
				log.Fatal(err)
			}
		}(ctx)
		tr := tp.Tracer("component-main")
		ctx, span := tr.Start(ctx, "foo")
		defer span.End()
		c.Next()
	}
}

func initTracer() (*trace.TracerProvider, error) {
	env := config.GetEnv()
	url := "http://jaeger:14268/api/traces"
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(env.App.Name),
			attribute.String("environment", env.App.Env),
			attribute.Int64("ID", 1),
		)),
		trace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
