package store

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
)

func (a *storeActor) delete( msg *message.Delete) core.Resolver {

  resolver := core.NewResolve()
  id, err := uuid.Parse(msg.Id)
  if err != nil {
    return resolver.Error(err)
  }

  return resolver.Mono(a.service.Delete(id)).Then(func(any types.Any) types.Any {
    return &core.Done{}
  })
}

