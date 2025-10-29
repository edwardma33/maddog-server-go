package logger

import (
	"log"
	"github.com/edwardma33/maddog-server-go"
)

func LoggerFilter(next maddog.Handler) maddog.Handler {
  return maddog.HandlerFunc(func(ctx *maddog.Context) {
    log.Printf("%s %s\n", ctx.Req.Method, ctx.Req.URL.Path)
    next.ServeHTTP(ctx)
  })
}
