package store

import (
  "errors"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (a *storeActor) fetch(msg *message.Fetch) core.Resolver {
  switch msg.By {
  case message.ID:
    return a.Resolve(a.query.SearchById(msg.Filter))

  case message.NAME:
   return a.Resolve(a.query.SearchByName(msg.Filter))

  case message.CODE:
    return a.Resolve(a.query.SearchByCode(msg.Filter))
  }

  resolve := core.NewResolve()
  return resolve.Error(errors.New("this message type is not implemented"))
}

func (a *storeActor) Resolve(maybe monad.Maybe) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Maybe(maybe).Then(func(any types.Any) types.Any {
    p, err := any.(monad.Mono).Unwrap()
    if err != nil {
      return err
    }
    return MapProductToMessage(p.(*entity.Product))
  })
}
