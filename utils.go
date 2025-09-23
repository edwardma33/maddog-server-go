package maddog


import (
	"context"
	"errors"
	"net/http"
	"strings"
)

func ParseDynamicPathVariables(r *http.Request, prefix string, name string) (string, error) {
  path := r.URL.Path

  if !strings.HasPrefix(path, prefix) {
    return "", errors.New("Not Found")
  }

  s := strings.TrimPrefix(path, prefix)

  ctx := context.WithValue(r.Context(), name, s)

  r = r.WithContext(ctx)

  return s, nil
}

type ResMap map[string]any
