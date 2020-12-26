package store

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/services"
  "github.com/mlambda-net/store/pkg/domain/utils"
)

type storeActor struct {
  service services.StoreService
  image   services.ImageService
  query   services.QueryService
}


func (a *storeActor) Receive(ctx actor.Context) {
  switch msg := ctx.Message().(type) {
  case *message.Add:
    ctx.Respond(a.add(msg).Response())

  case *message.CreateProduct:
    ctx.Respond(a.create(msg).Response())

  case *message.Fetch:
    ctx.Respond(a.fetch(msg).Response())

  case *message.EditProduct:
    ctx.Respond(a.edit(msg).Response())

  case *message.Search:
    ctx.Respond(a.search(msg).Response())

  case *message.Delete:
    ctx.Respond(a.delete(msg).Response())

  case *message.All:
    ctx.Respond(a.all(msg).Response())
  }
}


func NewStoreProps(config *utils.Configuration) *actor.Props {
  service := services.NewStoreService(config)
  image := services.NewImageService(config)
  query := services.NewQueryService(config)
  return actor.PropsFromProducer(func() actor.Actor {
    return &storeActor{
      service: service,
      query: query,
      image:   image,
    }
  })
}
