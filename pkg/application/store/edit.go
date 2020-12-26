package store

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (a *storeActor) edit(msg *message.EditProduct) core.Resolver {

  resolve := core.NewResolve()
  id, err := uuid.Parse(msg.Id)

  if err != nil {
    return resolve.Error(err)
  }

  var resolver core.Resolver
  a.service.EditProduct(&entity.Product{
    Name:        msg.Name,
    Description: msg.Description,
    ID:          id,
  }).Bind(func(any types.Any) monad.Maybe {
    resolver = resolve.Mono(any.(monad.Mono))
    return monad.Empty()
  })

  return resolver.Then(func(any types.Any) types.Any {
    return &core.Done{}
  })

}

