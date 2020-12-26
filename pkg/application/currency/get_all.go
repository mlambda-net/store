package currency

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (c currencyActor) GetAll(*message.CurrencyAll) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono( c.service.GetCurrencies()).Then(func(any types.Any) types.Any {
    if any != nil {
      currencies := any.([]*entity.Currency)
      items := make([]*message.Currency, 0)
      for _, cu := range currencies {
        items = append(items, &message.Currency{
          Name:      cu.Name,
          Symbol:    cu.Symbol,
          Character: cu.Character,
        })
      }
      return &message.Currencies{Currency: items}
    }
    return &message.Currencies{}
  })
}
