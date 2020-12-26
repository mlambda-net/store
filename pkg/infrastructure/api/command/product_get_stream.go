package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// CreateProduct godoc
// @Summary Get the product
// @Produce json
// @Success 200 {array} message.Product
// @Failure 500 {string} string "Internal error"
// @Router /product/stream [get]
func (c *control) GetProductStream(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Connection", "Keep-Alive")
  w.Header().Set("Transfer-Encoding", "chunked")
  w.Header().Set("X-Content-Type-Options", "nosniff")

  res, err := c.store.Request(&message.All{}).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  flusher, ok := w.(http.Flusher)
  if !ok {
    http.Error(w, "can not stream the response", http.StatusInternalServerError)
  } else {
    items := res.(*message.Products).Products
    for _, i := range  items {
      _ = json.NewEncoder(w).Encode(i)
      flusher.Flush()
    }
  }
}

