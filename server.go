package maddog

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type App struct {
  Server http.Server
  Router *chi.Mux
  UrlPrefix string
}

func NewServer(addr string, urlPrefix string) *App {
  r := chi.NewRouter()
  return &App{
    
    Server: http.Server{
      Handler: r,
      Addr: addr,
    },
    Router: r,
    UrlPrefix: urlPrefix,
  }
}

func (a *App) Run() error {
  fmt.Printf("Server running @ http://localhost%s\n", a.Server.Addr)
  return a.Server.ListenAndServe()
}

func (a *App) Get(pattern string, handler HandlerFunc) {
  a.Router.Get(pattern, Adapt(handler))
}

func (a *App) Post(pattern string, handler HandlerFunc) {
  a.Router.Post(pattern, Adapt(handler))
}

func (a *App) Put(pattern string, handler HandlerFunc) {
  a.Router.Put(pattern, Adapt(handler))
}

func (a *App) Delete(pattern string, handler HandlerFunc) {
  a.Router.Delete(pattern, Adapt(handler))
}

func (a *App) Patch(pattern string, handler HandlerFunc) {
  a.Router.Patch(pattern, Adapt(handler))
}

func (a *App) Handle(pattern string, handler HandlerFunc) {
  a.Router.Handle(fmt.Sprintf("%s%s", a.UrlPrefix, pattern), Adapt(handler))
}

func (a *App) HandleFs(pattern string, dir http.Dir) {
  a.Router.Handle(fmt.Sprintf("%s%s", a.UrlPrefix, pattern), http.StripPrefix(pattern, http.FileServer(dir)))
}

func (a *App) HandleTempl(pattern string, t *templ.ComponentHandler) {
  a.Router.Handle(fmt.Sprintf("%s%s", a.UrlPrefix, pattern), t)
}
