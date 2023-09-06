// Package middlewares – содержит middleware для обработки запросов
package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/arseniy96/url-shortener/internal/logger"
)

const (
	AcceptEncoding  = "Accept-Encoding"
	ContentEncoding = "Content-Encoding"
	EncodingType    = "gzip"
)

type gzipWriter struct {
	Writer  http.ResponseWriter
	ZWriter *gzip.Writer
}

func (w *gzipWriter) Write(b []byte) (int, error) {
	return w.ZWriter.Write(b)
}

func (w *gzipWriter) WriteHeader(statusCode int) {
	w.Writer.Header().Set(ContentEncoding, EncodingType)
	w.Writer.WriteHeader(statusCode)
}

func (w *gzipWriter) Header() http.Header {
	return w.Writer.Header()
}

func (w *gzipWriter) Close() {
	err := w.ZWriter.Close()
	if err != nil {
		logger.Log.Error(err)
	}
}

func newGzipWriter(w http.ResponseWriter) *gzipWriter {
	return &gzipWriter{
		Writer:  w,
		ZWriter: gzip.NewWriter(w),
	}
}

type gzipReader struct {
	Reader  io.ReadCloser
	ZReader *gzip.Reader
}

func (r *gzipReader) Close() error {
	if err := r.Reader.Close(); err != nil {
		return err
	}
	return r.ZReader.Close()
}

func (r *gzipReader) Read(p []byte) (n int, err error) {
	return r.ZReader.Read(p)
}

func newGzipReader(reader io.ReadCloser) (*gzipReader, error) {
	newReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	return &gzipReader{
		Reader:  reader,
		ZReader: newReader,
	}, nil
}

// GzipMiddleware – миддлваря для сжатия.
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentWriter := w

		if strings.Contains(r.URL.Path, "pprof") {
			next.ServeHTTP(currentWriter, r)
			return
		}

		if strings.Contains(r.Header.Get(AcceptEncoding), EncodingType) {
			newWriter := newGzipWriter(w)
			currentWriter = newWriter
			defer newWriter.Close()
		}

		if strings.Contains(r.Header.Get(ContentEncoding), EncodingType) {
			reader, err := newGzipReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			r.Body = reader
			defer func() {
				if err := reader.Close(); err != nil {
					logger.Log.Error(err)
				}
			}()
		}

		next.ServeHTTP(currentWriter, r)
	})
}
