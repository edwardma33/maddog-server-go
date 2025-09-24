package maddog


type Filter func(Handler) Handler

func FilterChain(filters ...Filter) Filter {
  return func(finalHandler Handler) Handler {
    h := finalHandler

    for i := len(filters) - 1; i>=0; i-- {
      h = filters[i](h)
    }
    return h
  }
}
