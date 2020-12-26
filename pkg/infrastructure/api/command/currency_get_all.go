package command

import (
  "encoding/json"
  "github.com/mlambda-net/store/pkg/application/message"
  "net/http"
)

// CurrencyAll godoc
// @Summary Get the currencies
// @Produce json
// @Success 200 {array} message.Currency
// @Failure 500 {string} string "Internal error"
// @Router /currency [get]
func (c *control) CurrencyAll(w http.ResponseWriter, r *http.Request) {
  res, err := c.currency.Request(&message.CurrencyAll{}).Unwrap()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  items := res.(*message.Currencies).Currency
  _ = json.NewEncoder(w).Encode(items)
}

