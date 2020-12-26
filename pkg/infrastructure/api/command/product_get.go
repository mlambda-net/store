package command

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
  "strings"
)

// GetProduct godoc
// @Summary Get the product
// @Produce json
// @Param id path string false "search by id"
// @Success 200 {object} message.Product
// @Failure 404 {string} string "Not found"
// @Failure 500 {string} string "Internal error"
// @Router /product/{id} [get]
func (c *control) GetProduct(w http.ResponseWriter, r *http.Request) {
  data := mux.Vars(r)["id"]
  res, err := c.store.Request(&message.Fetch{
    Filter: data,
    By:     message.ID,
  }).Unwrap()
  if err != nil {
     msg := err.Error()
     if strings.Contains(msg, "404") {
       http.Error(w, err.Error(), http.StatusNotFound)
     } else {
       http.Error(w, err.Error(), http.StatusInternalServerError)
     }

  } else {
    _ = json.NewEncoder(w).Encode(res)
  }
}
