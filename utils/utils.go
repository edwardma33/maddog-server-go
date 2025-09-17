package utils


import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)


func MapReqBody(r *http.Request, v any) any {
  data, _ := io.ReadAll(r.Body)
  json.Unmarshal(data, &v)
  return v
}

func WriteJson(w http.ResponseWriter, status int, v any) {
  data, _ := json.Marshal(v)
  w.WriteHeader(status)
  w.Write(data)
}

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
