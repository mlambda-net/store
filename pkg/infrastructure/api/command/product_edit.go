package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// EditProduct godoc
// @Summary Edit the product
// @Produce json
// @Param data body message.EditProduct true "edit product"
// @Success 200 {object} core.Done
// @Failure 500 {string} string "Internal error"
// @Router /product [put]
func (c *control) EditProduct(w http.ResponseWriter, r *http.Request) {

  var data message.EditProduct
  _ = json.NewDecoder(r.Body).Decode(&data)
  id, err := c.store.Request(&data).Unwrap()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
