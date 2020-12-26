package command

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// CategoryAll godoc
// @Summary Get the categories
// @Produce json
// @Param lang path string false "get categories"
// @Success 200 {array} message.Lookup
// @Failure 500 {string} string "Internal error"
// @Router /category/{lang} [get]
func (c *control) CategoryAll(w http.ResponseWriter, r *http.Request) {
  lang := mux.Vars(r)["lang"]
  res, err := c.lookup.Request(&message.Categories{Lang: lang}).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  items := res.(*message.Lookups).Lookups
  _ = json.NewEncoder(w).Encode(items)
}

