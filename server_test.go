package maddog_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwardma33/maddog-server-go"
)

func TestPing(t *testing.T) {
	srv := maddog.NewServer(":8080", "/")

	srv.Get("/ping", func(ctx *maddog.Context) {
		ctx.Res.Write([]byte("pong"))
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	srv.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, go %d", w.Code)
	}
}

func TestUse(t *testing.T) {
	srv := maddog.NewServer(":8080", "/")

	srv.Use(func(h maddog.Handler) maddog.Handler {
    return maddog.HandlerFunc(func(ctx *maddog.Context) {
      c := context.WithValue(ctx.Req.Context(), "msg", "COYG!")
      ctx.Req = ctx.Req.WithContext(c)

      h.ServeHTTP(ctx)
    })
  })

	srv.Get("/ping", func(ctx *maddog.Context) {
		t.Logf("Middleware msg: %s", ctx.Req.Context().Value("msg"))
		ctx.Res.Write([]byte("pong"))
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	srv.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, go %d", w.Code)
	}
}

func TestSubRouter(t *testing.T) {
	srv := maddog.NewServer(":8080", "/")

  userRouter := srv.NewSubRouterMount("/users")
  
  userRouter.Get("/create", func(ctx *maddog.Context) {
    res := maddog.ResMap{"msg": "User created successfully"}
    t.Log(res)
    ctx.WriteJSON(200, res)
  })

	req := httptest.NewRequest(http.MethodGet, "/users/create", nil)
	w := httptest.NewRecorder()

	srv.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, go %d", w.Code)
	}
}
