package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
  http.ResponseWriter;
  statusCode int;
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
  w.ResponseWriter.WriteHeader(statusCode)
  w.statusCode = statusCode
}

func Logging(handler http.Handler) http.Handler {
  return http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
    start := time.Now()
    log.Println(request.Method, request.URL.Path, "incomming request")
    ww := &wrappedWriter{
      ResponseWriter: writer,
      statusCode: http.StatusOK,
    }
    handler.ServeHTTP(ww, request)
    log.Println(request.Method, request.URL.Path, ww.statusCode, time.Since(start))
  })
}
