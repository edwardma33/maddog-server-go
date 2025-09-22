package maddog

import (
	"net/http"
)

type Controller struct {
  UrlPrefix string
  Routes []Route
}

type Route struct {
  Path string
  Handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
  return Route{
    Path: path,
    Handler: handler,
  }
}
