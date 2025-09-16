package cors

import (
	"log"
	"net/http"
)

func DefaultCorsFilter(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Cors Filter Hit")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.Header().Set("Access-Control-Expose-Headers", "Authorization")
    
    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }
    
    next.ServeHTTP(w, r)
  })
}

type CorsConfig struct {
  Origins string
  Methods string
  Headers string
}

func CustomCorsFilter(next http.Handler, corsCfg CorsConfig) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("Cors Filter Hit")
    w.Header().Set("Access-Control-Allow-Origin", corsCfg.Origins)
    w.Header().Set("Access-Control-Allow-Methods", corsCfg.Methods)
    w.Header().Set("Access-Control-Allow-Headers", corsCfg.Headers)
    w.Header().Set("Access-Control-Expose-Headers", "Authorization")
    
    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }
    
    next.ServeHTTP(w, r)
  })
}
