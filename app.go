package maddog

import (
	"fmt"
	"net/http"
	"github.com/edwardma33/maddog-server-go/jwt"
	"github.com/gorilla/sessions"
)

type App struct {
  Server http.Server
  Store *sessions.CookieStore
  Cfg AppConfig
}

func NewApp(addr string, cfg AppConfig) *App {
  return &App{
    Server: http.Server{
      Handler: cfg.Mux,
      Addr: addr,
    },
    Store: sessions.NewCookieStore(cfg.SessionKey),
    Cfg: cfg,
  }
}

func (a *App) Run() error {
  fmt.Printf("Server running @ http://localhost%s\n", a.Server.Addr)
  return a.Server.ListenAndServe()
}

func (a *App) HandleCustom(pattern string, handler HandlerFunc) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), Adapt(handler))
}

func (a *App) Handle(pattern string, handler HandlerFunc) {
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), a.Cfg.GlobalFilterChain(Adapt(handler)))
}

func (a *App) HandleProtected(pattern string, handler HandlerFunc) {
  filterChain := FilterChain(jwt.DefaultJwtFilter, a.Cfg.GlobalFilterChain)
  a.Cfg.Mux.Handle(fmt.Sprintf("%s%s", a.Cfg.UrlPrefix, pattern), filterChain(Adapt(handler)))
}

func (a *App) RegisterControllers(controllers []Controller) {
  for _, c := range controllers {
    for _, r := range c.Routes {
      a.Handle(fmt.Sprintf("%s%s", c.UrlPrefix, r.Path), r.Handler)
    }
  }
}

type AppConfig struct {
  Mux *http.ServeMux
  UrlPrefix string
  GlobalFilterChain Filter
  SessionKey []byte
}

func NewAppConfig(urlPrefix, sessionKey string, globalFilters []Filter) AppConfig {
  sessionKeyBytes := []byte(sessionKey)
  return AppConfig{
    Mux: http.NewServeMux(),
    UrlPrefix: urlPrefix,
    SessionKey: sessionKeyBytes,
    GlobalFilterChain: FilterChain(globalFilters...),
  }
}
