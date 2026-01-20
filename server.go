package maddog

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)



type App struct {
  Server http.Server
  MainRouter *chi.Mux
  Router *chi.Mux
  GlobalPattern string
}

func NewServer(addr string, globalPattern string) *App {
  r := chi.NewRouter()
  sr := &SubRouter{
    Pattern: globalPattern,
    Router: chi.NewRouter(),
  }
  return &App{
    Server: http.Server{
      Handler: r,
      Addr: addr,
    },
    GlobalPattern: globalPattern,
    MainRouter: r,
    Router: sr.Router,
  }
}

func (a *App) Run() error {
  a.MainRouter.Mount(a.GlobalPattern, a.Router)
  fmt.Printf("Server running @ http://localhost%s\n", a.Server.Addr)
  return a.Server.ListenAndServe()
}

func (a *App) Use(middlewares ...func(Handler) Handler) {
  adapted := []func(http.Handler) http.Handler{}
  for _, mw := range middlewares {
    adapted = append(adapted, AdaptMiddleware(mw))
  }
  a.Router.Use(adapted...)
}

func (a *App) NewSubRouter(pattern string) *SubRouter {
  return &SubRouter{
    Pattern: pattern,
    Router: chi.NewRouter(),
  }
}

func (a *App) NewSubRouterMount(pattern string) *SubRouter {
  sr := &SubRouter{
    Pattern: pattern,
    Router: chi.NewRouter(),
  }
  a.Mount(sr)
  return sr
}

func (a *App) Mount(sr *SubRouter) {
  a.Router.Mount(sr.Pattern, sr.Router)
}

func (a *App) Get(pattern string, handler HandlerFunc) {
  a.Router.Get(pattern, AdaptF(handler))
}

func (a *App) Post(pattern string, handler HandlerFunc) {
  a.Router.Post(pattern, AdaptF(handler))
}

func (a *App) Put(pattern string, handler HandlerFunc) {
  a.Router.Put(pattern, AdaptF(handler))
}

func (a *App) Delete(pattern string, handler HandlerFunc) {
  a.Router.Delete(pattern, AdaptF(handler))
}

func (a *App) Patch(pattern string, handler HandlerFunc) {
  a.Router.Patch(pattern, AdaptF(handler))
}

func (a *App) HandleFs(pattern string, dir http.Dir) {
  a.Router.Handle(pattern, http.StripPrefix(pattern, http.FileServer(dir)))
}

func (a *App) HandleTempl(pattern string, t *templ.ComponentHandler) {
  a.Router.Handle(pattern, t)
}
