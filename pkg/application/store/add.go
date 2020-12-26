package store

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (a *storeActor) add(msg *message.Add) core.Resolver {
  resolve := core.NewResolve()
  return resolve.Mono(a.service.Add(&entity.Product{
    ID:          uuid.New(),
    Name:        msg.Name,
    Description: msg.Description,
    Price:       msg.Price,
    Code:        msg.Code,
  })).Then(func(any types.Any) types.Any {
    return &message.ProductId{Id: any.(*entity.Product).ID.String()}
  })
}

