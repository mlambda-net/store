package lookup

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
)

func (c Actor) CategoryAdd(ms *message.CategoryAdd) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(c.service.AddCategory(ms.Name)).Then(func(any types.Any) types.Any {
    return message.Done{}
  })
}
