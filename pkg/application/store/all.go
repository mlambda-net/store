package store

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (a *storeActor) all(msg *message.All)  core.Resolver {
  resolve := core.NewResolve()
  return resolve.Maybes(a.query.GetAll()).Then(func(any types.Any) types.Any {
    switch ms := any.(type) {
    case *core.Empty:
      return &message.Products{}
    case []interface{}:
      var items []*message.Product
      for _, m := range ms {
        m.(monad.Mono).Bind(func(any types.Any) monad.Mono {
          prod := any.(*entity.Product)
          items = append(items, MapProductToMessage(prod))
          return monad.ToMono(prod)
        })
      }
      return &message.Products{Products: items}
    default:
      return &message.Products{}
    }
  })
}

