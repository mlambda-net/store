package command

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// EditProduct godoc
// @Summary Delete the product
// @Produce json
// @Param id path string false "search by id"
// @Success 200 {object} message.Delete
// @Failure 500 {string} string "Internal error"
// @Router /product/{id} [delete]
func (c *control) DeleteProduct(w http.ResponseWriter, r *http.Request) {
  data := mux.Vars(r)["id"]
  id, err := c.store.Request(&message.Delete{Id: data}).Unwrap()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(id)
  }
}
