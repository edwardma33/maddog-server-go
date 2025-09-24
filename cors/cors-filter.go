package cors

import (
	"log"
	"net/http"
	"strconv"
	"github.com/edwardma33/maddog-server-go"
)

const (
  AllMethods = "GET, POST, PUT, DELETE, OPTIONS"
)

func DefaultCorsFilter(next maddog.Handler) maddog.Handler {
  return maddog.HandlerFunc(func(ctx *maddog.Context) {
    log.Println("Cors Filter Hit")
    ctx.Res.Header().Set("Access-Control-Allow-Origin", "*")
    ctx.Res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    ctx.Res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    ctx.Res.Header().Set("Access-Control-Expose-Headers", "Authorization")
    ctx.Res.Header().Set("Access-Control-Allow-Credentials", "true")
    
    if ctx.Req.Method == http.MethodOptions {
      ctx.Res.WriteHeader(http.StatusNoContent)
      return
    }
    
    next.ServeHTTP(ctx)
  })
}

type CorsConfig struct {
  Origins string
  Methods string
  Headers string
  AllowCredentials bool
}

func CustomCorsFilter(cfg CorsConfig) maddog.Filter {
  return func(next maddog.Handler) maddog.Handler {
    return maddog.HandlerFunc(func(ctx *maddog.Context) {
    log.Println("Cors Filter Hit")

    ctx.Res.Header().Set("Access-Control-Allow-Origin", cfg.Origins)
    ctx.Res.Header().Set("Access-Control-Allow-Methods", cfg.Methods)
    ctx.Res.Header().Set("Access-Control-Allow-Headers", cfg.Headers)
    ctx.Res.Header().Set("Access-Control-Expose-Headers", "Authorization")
    ctx.Res.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(cfg.AllowCredentials))
    
    if ctx.Req.Method == http.MethodOptions {
      ctx.Res.WriteHeader(http.StatusNoContent)
      return
    }
    
    next.ServeHTTP(ctx)
    })
  }
}
