package maddog

import "net/http"

type Filter func(http.Handler) http.Handler

func FilterChain(filters ...Filter) Filter {
  return func(finalHandler http.Handler) http.Handler {
    h := finalHandler

    for i := len(filters) - 1; i>=0; i-- {
      h = filters[i](h)
    }
    return h
  }
}
