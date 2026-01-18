package maddog

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
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

func (a *App) Handle(pattern string, handler HandlerFunc, filters ...Filter) {
  filterChain := FilterChain(append(a.Cfg.GlobalFilters, filters...)...)
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), AdaptH(filterChain(HandlerFunc(handler))))
}

func (a *App) HandleFs(pattern string, dir http.Dir) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), http.StripPrefix(pattern, http.FileServer(dir)))
}

func (a *App) HandleTempl(pattern string, t *templ.ComponentHandler) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), t)
}

func (a *App) RegisterControllers(controllers []Controller) {
  for _, c := range controllers {
    for _, r := range c.Routes {
      a.Handle(fmt.Sprintf("%s%s", c.UrlPrefix, r.Path), r.Handler, r.Filters...)
    }
  }
}

type AppConfig struct {
  Mux *chi.Mux
  UrlPrefix string
  GlobalFilters []Filter
  SessionKey []byte
}

func NewAppConfig(urlPrefix, sessionKey string, globalFilters []Filter) AppConfig {
  sessionKeyBytes := []byte(sessionKey)
  return AppConfig{
    Mux: chi.NewRouter(),
    UrlPrefix: urlPrefix,
    SessionKey: sessionKeyBytes,
    GlobalFilters: globalFilters,
  }
}
