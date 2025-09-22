package maddog

type Controller struct {
  UrlPrefix string
  Routes []Route
}

type Route struct {
  Path string
  Handler HandlerFunc
}

func NewRoute(path string, handler HandlerFunc) Route {
  return Route{
    Path: path,
    Handler: handler,
  }
}
