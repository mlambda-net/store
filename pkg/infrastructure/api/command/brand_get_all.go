package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// BrandAll godoc
// @Summary Get the brands
// @Produce json
// @Success 200 {array} message.Lookup
// @Failure 500 {string} string "Internal error"
// @Router /brand [get]
func (c *control) BrandAll(w http.ResponseWriter, r *http.Request) {
  res, err := c.lookup.Request(&message.Brands{}).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  items := res.(*message.Lookups).Lookups
  _ = json.NewEncoder(w).Encode(items)
}

