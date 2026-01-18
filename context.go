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

func (c *Context) Redirect(url string, status int) {
	http.Redirect(c.Res, c.Req, url, status)
}

func (c *Context) Error(status int, err error) {
	res := ResMap{"error": err.Error()}
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

func (c *Context) Cookie(name string) (*http.Cookie, error) {
	return c.Req.Cookie(name)
}

func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.Res, cookie)
}

func (c *Context) DeleteCookie(name string) {
	http.SetCookie(c.Res, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

type Handler interface {
	ServeHTTP(*Context)
}

type HandlerFunc func(*Context)

// Make HandlerFunc implement Handler
func (f HandlerFunc) ServeHTTP(c *Context) {
	f(c)
}

func AdaptH(h HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{Res: w, Req: r}
		h.ServeHTTP(ctx)
	})
}

func AdaptF(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{Res: w, Req: r}
		h.ServeHTTP(ctx)
	}
}

type httpHandlerAdapter struct {
	handler http.Handler
}

func (a httpHandlerAdapter) ServeHTTP(c *Context) {
	a.handler.ServeHTTP(c.Res, c.Req)
}

func AdaptMiddleware(mw func(Handler) Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		adapted := httpHandlerAdapter{next}
		wrapped := mw(adapted)
		return AdaptH(wrapped.(HandlerFunc))
	}
}
