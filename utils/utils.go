package utils


import (
	"encoding/json"
	"io"
	"net/http"
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
