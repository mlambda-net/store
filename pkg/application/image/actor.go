package image

import (
  "encoding/base64"
  "github.com/AsynkronIT/protoactor-go/actor"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/services"
  "github.com/mlambda-net/store/pkg/domain/utils"
)

type Actor struct {
  service services.ImageService
}

func (a Actor) Receive(ctx actor.Context) {
  switch ms := ctx.Message().(type) {
  case *message.ImageAdd:
    ctx.Respond(a.AddImage(ms).Response())
  }
}

func (a Actor) AddImage(ms *message.ImageAdd) core.Resolver {
  resolve := core.NewResolve()
  key := base64.RawURLEncoding.EncodeToString([]byte(ms.Id))
  var file = &entity.File{
    Name:    ms.File,
    Prefix:  key,
    Content: ms.Content,
  }
  return resolve.Mono(a.service.AddImage(file)).Then(func(any types.Any) types.Any {
    return &message.ImageUrl{
      Url: any.(*entity.Image).Path,
    }
  })
}

func NewImageActor(config *utils.Configuration) *actor.Props {
  service := services.NewImageService(config)
  return actor.PropsFromProducer(func() actor.Actor {
    return &Actor{
      service: service,
    }
  })
}
