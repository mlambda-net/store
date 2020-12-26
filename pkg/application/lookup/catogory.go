package lookup

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (c Actor) Categories(m *message.Categories) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(c.query.Categories()).Then(func(any types.Any) types.Any {
    if any != nil {
      items := any.([]*entity.Lookup)
      data := make([]*message.Lookup, 0)
      for _, i := range items {
        data = append(data, &message.Lookup{Name: i.Name})
      }
      return &message.Lookups{Lookups: data}
    }
    return &message.Lookups{}
  })
}

