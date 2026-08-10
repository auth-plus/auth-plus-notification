package middlewares

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Trace is a middleware that gather trace of system
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tr := otel.Tracer("middlewares.Trace")
		
		// Extract context from request to support distributed tracing propagation
		ctx := c.Request.Context()
		ctx, span := tr.Start(ctx, "main")
		defer span.End()

		span.SetAttributes(
			attribute.String("HTTP_URL", c.Request.URL.Path),
			attribute.String("HTTP_METHOD", c.Request.Method),
		)
		span.AddEvent("HTTP_STARTED")

		// Safely read and restore request body
		var requestBody []byte
		if c.Request.Body != nil {
			var err error
			requestBody, err = c.GetRawData()
			if err != nil {
				log.Println("Error reading request body:", err)
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// Wrap response writer to capture the response body
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// Update request with the new trace context
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		span.AddEvent("HTTP_FINISHED")

		span.SetAttributes(
			attribute.Int("HTTP_STATUS_CODE", c.Writer.Status()),
			attribute.String("response.body", blw.body.String()),
			attribute.String("request.body", string(requestBody)),
		)
	}
}
