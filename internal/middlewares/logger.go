package middlewares

import (
	"net/http"
	"time"

	"github.com/arseniy96/url-shortener/internal/logger"
)

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lw.ResponseWriter.Write(b)
	lw.responseData.size = size
	return size, err
}

func (lw *loggingResponseWriter) WriteHeader(statusCode int) {
	lw.ResponseWriter.WriteHeader(statusCode)
	lw.responseData.status = statusCode
}

// LoggerMiddleware – миддлваря для логирования запросов
// Логируем метод запроса, путь, время выполнения, статус и размер ответа
func LoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   &responseData{},
		}

		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		logger.Log.Infow(
			"HTTP Request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", duration,
			"response_status", lw.responseData.status,
			"response_size", lw.responseData.size,
		)
	})
}
