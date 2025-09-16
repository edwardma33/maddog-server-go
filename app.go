package maddog

import (
	"fmt"
	"net/http"

	"github.com/edwardma33/maddog-server-go/jwt"
)

type App struct {
  Server http.Server
  Cfg AppConfig
}

func NewApp(addr string, cfg AppConfig) *App {
  return &App{
    Server: http.Server{
      Handler: cfg.Mux,
      Addr: addr,
    },
    Cfg: cfg,
  }
}

func (a *App) Run() error {
  fmt.Printf("Server running @ http://localhost%s\n", a.Server.Addr)
  return a.Server.ListenAndServe()
}

func (a *App) HandleCustom(pattern string, handler http.Handler) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), handler)
}

func (a *App) Handle(pattern string, handler http.Handler) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), a.Cfg.GlobalFilterChain(handler))
}

func (a *App) HandleProtected(pattern string, handler http.Handler) {
  filterChain := FilterChain(jwt.DefaultJwtFilter, a.Cfg.GlobalFilterChain)
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), filterChain(handler))
}

type AppConfig struct {
  Mux *http.ServeMux
  UrlPrefix string
  GlobalFilterChain Filter
}

func NewAppConfig(urlPrefix string, globalFilters []Filter) AppConfig {
  return AppConfig{
    Mux: http.NewServeMux(),
    UrlPrefix: urlPrefix,
    GlobalFilterChain: FilterChain(globalFilters...),
  }
}

type Filter func(http.Handler) http.Handler

func FilterChain(filters ...Filter) Filter {
  return func(finalHandler http.Handler) http.Handler {
    h := finalHandler

    for i := len(filters) - 1; i>=0; i-- {
      h = filters[i](h)
    }
    return h
  }
}
