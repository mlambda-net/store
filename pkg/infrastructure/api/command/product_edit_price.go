package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// EditProduct godoc
// @Summary edit product price
// @Produce json
// @Param data body message.EditPrice true "edit product price"
// @Success 200 {object} core.Done
// @Failure 500 {string} string "Internal error"
// @Router /product/price [put]
func (c *control) EditPrice(w http.ResponseWriter, r *http.Request) {

  var data message.EditProduct
  _ = json.NewDecoder(r.Body).Decode(&data)
  id, err := c.store.Request(&data).Unwrap()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
