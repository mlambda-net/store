package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// GetProductAll godoc
// @Summary Get the products
// @Produce json
// @Success 200 {array} message.Product
// @Failure 500 {string} string "Internal error"
// @Router /product [get]
func (c *control) GetProductAll(w http.ResponseWriter, r *http.Request) {
  res, err := c.store.Request(&message.All{}).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  items := res.(*message.Products).Products
  _ = json.NewEncoder(w).Encode(items)
}

