package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// CategoryAdd godoc
// @Summary Add the category
// @Produce json
// @Param data body message.CategoryAdd true "create category"
// @Success 200 {object} message.Done
// @Failure 500 {string} string "Internal error"
// @Router /category [post]
func (c *control) CategoryAdd(w http.ResponseWriter, r *http.Request) {
  var data message.CategoryAdd
  _ = json.NewDecoder(r.Body).Decode(&data)
  rsp, err := c.lookup.Request(&data).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  } else {
    _ = json.NewEncoder(w).Encode(rsp)
  }
}

