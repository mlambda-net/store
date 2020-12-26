package currency

import (
  "github.com/AsynkronIT/protoactor-go/actor"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/services"
  "github.com/mlambda-net/store/pkg/domain/utils"
)

type currencyActor struct {
  service services.CurrencyService
}

func (c currencyActor) Receive(ctx actor.Context) {
  switch ms := ctx.Message().(type) {
  case *message.CurrencyAll:
    ctx.Respond(c.GetAll(ms).Response())
  }
}


func NewCurrencyActor(config *utils.Configuration)  *actor.Props {
  service := services.NewCurrencyService(config)
  return actor.PropsFromProducer(func() actor.Actor {
    return &currencyActor {
      service: service,
    }
  })
}
