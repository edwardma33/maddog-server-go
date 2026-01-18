package maddog_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwardma33/maddog-server-go"
)

func TestPing(t *testing.T) {
  srv := maddog.NewServer(":8080", "")

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
