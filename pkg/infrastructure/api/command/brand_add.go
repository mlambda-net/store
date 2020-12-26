package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// BrandAdd godoc
// @Summary Add the brand
// @Produce json
// @Param data body message.BrandAdd true "create brand"
// @Success 200 {object} message.Done
// @Failure 500 {string} string "Internal error"
// @Router /brand [post]
func (c *control) BrandAdd(w http.ResponseWriter, r *http.Request) {
  var data message.BrandAdd
  _ = json.NewDecoder(r.Body).Decode(&data)
  rsp, err := c.lookup.Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(rsp)
  }
}
