package maddog

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
  Res http.ResponseWriter
  Req *http.Request
}

func (c *Context) WriteJSON(status int, res ResMap) {
  data, _ := json.Marshal(res)
  c.Res.WriteHeader(status)
  c.Res.Write(data)
}

func (c *Context) MapReqJSON(v any) (any, error) {
  data, err := io.ReadAll(c.Req.Body)
  if err != nil {
    return nil, err
  }
  defer c.Req.Body.Close()
  err = json.Unmarshal(data, &v)
  return v, err
}

type Handler interface {
  ServeHTTP(*Context)
}

type HandlerFunc func(*Context)

// Make HandlerFunc implement Handler
func (f HandlerFunc) ServeHTTP(c *Context) {
  f(c)
}

func Adapt(h HandlerFunc) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := &Context{Res: w, Req: r}
    h.ServeHTTP(ctx)
  })
}

