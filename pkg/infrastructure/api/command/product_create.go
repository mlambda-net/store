package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// CreateProduct godoc
// @Summary Create the product
// @Produce json
// @Param data body message.CreateProduct true "create product"
// @Success 200 {object} message.ProductId
// @Failure 500 {string} string "Internal error"
// @Router /product [post]
func (c *control) CreateProduct(w http.ResponseWriter, r *http.Request) {
  var data message.CreateProduct
  _ = json.NewDecoder(r.Body).Decode(&data)
  id, err := c.store.Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
