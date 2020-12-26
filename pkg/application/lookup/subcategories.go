package lookup

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (c Actor) SubCategories(category string) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Maybe(c.service.SubCategories(category)).Then(func(any types.Any) types.Any {
    items, err := any.(monad.Mono).Unwrap()
    if err != nil {
      return err
    }
    ls := make([]*message.Lookup,0)
    for _, i := range items.([]*entity.Lookup) {
      ls = append(ls, &message.Lookup{
        Name: i.Name,
      })
    }
    return &message.Lookups{Lookups: ls}
  })
}

