package lookup

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
)

func (c Actor) BrandAdd(ms *message.BrandAdd) core.Resolver{
  resolve := core.NewResolve()
  return resolve.Mono(c.service.AddBrand(ms.Name)).Then(func(any types.Any) types.Any {
    return message.Done{}
  })
}
