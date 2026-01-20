package maddog

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type SubRouter struct {
  Pattern string
  Router *chi.Mux
}

func (a *SubRouter) Use(middlewares ...func(Handler) Handler) {
  adapted := []func(http.Handler) http.Handler{}
  for _, mw := range middlewares {
    adapted = append(adapted, AdaptMiddleware(mw))
  }
  a.Router.Use(adapted...)
}

func (r *SubRouter) Get(pattern string, handler HandlerFunc) {
  r.Router.Get(pattern, AdaptF(handler))
}

func (r *SubRouter) Post(pattern string, handler HandlerFunc) {
  r.Router.Post(pattern, AdaptF(handler))
}

func (r *SubRouter) Put(pattern string, handler HandlerFunc) {
  r.Router.Put(pattern, AdaptF(handler))
}

func (r *SubRouter) Delete(pattern string, handler HandlerFunc) {
  r.Router.Delete(pattern, AdaptF(handler))
}

func (r *SubRouter) Patch(pattern string, handler HandlerFunc) {
  r.Router.Patch(pattern, AdaptF(handler))
}

func (r *SubRouter) HandleFs(pattern string, dir http.Dir) {
  r.Router.Handle(pattern, http.StripPrefix(pattern, http.FileServer(dir)))
}

func (r *SubRouter) HandleTempl(pattern string, t *templ.ComponentHandler) {
  r.Router.Handle(pattern, t)
}
