package maddog

type Controller struct {
  UrlPrefix string
  Routes []Route
}

type Route struct {
  Path string
  Handler HandlerFunc
  Filters []Filter
}

func NewRoute(path string, handler HandlerFunc, filters ...Filter) Route {
  return Route{
    Path: path,
    Handler: handler,
    Filters: filters,
  }
}

