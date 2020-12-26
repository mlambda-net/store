package lookup

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/services"
  "github.com/mlambda-net/store/pkg/domain/utils"
)

type Actor struct {
  service   services.LookupService
  query     services.LookupQuery
  translate services.TranslateService
}

func (c Actor) Receive(ctx actor.Context) {
  switch ms := ctx.Message().(type) {
  case *message.Categories:
    ctx.Respond(c.Categories(ms).Response())
  case *message.CategoryAdd:
    ctx.Respond(c.CategoryAdd(ms).Response())
  case *message.SubCategory:
    ctx.Respond(c.SubCategories(ms.Category).Response())
  case *message.Brands:
    ctx.Respond(c.Brands(ms).Response())
  case *message.BrandAdd:
    ctx.Respond(c.BrandAdd(ms).Response())
  }
}

func NewLookupActor(config *utils.Configuration) *actor.Props {
  service := services.NewLookupService(config)
  query := services.NewLookupQuery(config)
  return actor.PropsFromProducer(func() actor.Actor {
    return &Actor{
      service: service,
      query: query,
    }
  })
}
