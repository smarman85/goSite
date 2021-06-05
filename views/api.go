package views

import (
  "fmt"
  "net/http"
)

func ApiRender(w http.ResponseWriter, payload []byte) {
  fmt.Fprintf(w, string(payload))
}
