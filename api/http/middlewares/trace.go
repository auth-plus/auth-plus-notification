// Package middlewares contains all middleware for GIN
package middlewares

import (
	"auth-plus-notification/config"
	"bytes"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Trace is a middleware that gather trace of system
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		env := config.GetEnv()
		url := "http://zipkin:9411/api/v2/spans"
		exporter, exporterErr := zipkin.New(url)
		if exporterErr != nil {
			log.Fatal(exporterErr)
		}
		tp := trace.NewTracerProvider(
			trace.WithResource(resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(env.App.Name),
				attribute.String("environment", env.App.Env),
				attribute.String("HTTP_URL", c.Request.URL.Path),
				attribute.String("HTTP_METHOD", c.Request.Method),
			)),
			trace.WithBatcher(exporter),
		)
		otel.SetTracerProvider(tp)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		))
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
		tr := tp.Tracer("middlewares.Trace")
		ctx, span := tr.Start(ctx, "main")
		defer span.End()
		span.AddEvent("HTTP_STARTED")
		c.Next()
		span.AddEvent("HTTP_FINISHED")
		requestBody, err := c.GetRawData()
		if err != nil {
			log.Println(err)
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		span.SetAttributes(
			attribute.Int("HTTP_STATUS_CODE", c.Writer.Status()),
			attribute.String("response.body", string(requestBody)),
			attribute.String("request.body", blw.body.String()),
		)

	}
}
