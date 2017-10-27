package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggingResponsewriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponsewriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *loggingResponsewriter) Write(data []byte) (int, error) {
	log.Println("writing something")
	return lrw.ResponseWriter.Write(data)
}

type Logger struct {
	handler http.Handler
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//assume status is valid until proven otherwise
	lrw := &loggingResponsewriter{w, http.StatusOK}
	start := time.Now()
	l.handler.ServeHTTP(lrw, r)
	log.Printf("%s %s %d %v", r.Method, r.URL.Path, lrw.statusCode, time.Since(start))
}
