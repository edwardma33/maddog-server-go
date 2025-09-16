package logger

import (
	"log"
	"net/http"
)

func LoggerFilter(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    next.ServeHTTP(w, r)
    log.Printf("%s %s\n", r.Method, r.URL.Path)
  })
}
