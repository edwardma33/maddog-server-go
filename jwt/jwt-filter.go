package jwt

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/edwardma33/maddog-server-go/utils"
)

// used for protected routes
func DefaultJwtFilter(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    resMap := make(map[string]any)
    log.Println("Jwt Filter Hit")
    
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" || !strings.Contains(authHeader, "Bearer ") {
      resMap["error"] = "Missing or invalid jwt"
      utils.WriteJson(w, http.StatusUnauthorized, resMap)
      return
    }

    jws := authHeader[7:]

    sub, err := ValidateToken(jws)
    if err != nil {
      resMap["error"] = err.Error()
      utils.WriteJson(w, http.StatusUnauthorized, resMap)
      return
    }
    id, _ := strconv.Atoi(sub)
    ctx := context.WithValue(r.Context(), "id", id)

    next.ServeHTTP(w, r.WithContext(ctx))
  })
}
