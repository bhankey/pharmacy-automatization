package http

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/pborman/uuid"
	"github.com/sirupsen/logrus"
)

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func LoggingMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		requestID := uuid.NewUUID().String()
		f := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.WithFields(logrus.Fields{
						"err":        err,
						"trace":      debug.Stack(),
						"request_id": requestID,
					},
					).Error(
						"panic",
					)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			log := logger.WithFields(logrus.Fields{
				"status":     wrapped.status,
				"method":     r.Method,
				"path":       r.URL.EscapedPath(),
				"duration":   time.Since(start),
				"request_id": requestID,
			})

			if wrapped.Status() != http.StatusOK {
				log.Error("request")
			} else {
				log.Info("request")
			}
		}

		return http.HandlerFunc(f)
	}
}
