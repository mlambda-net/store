package store

import (
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/sirupsen/logrus"
)

func (a *storeActor) search(msg *message.Search) core.Resolver {
  resolve := core.NewResolve()

  return resolve.Maybes(a.query.Search(msg.Value)).Then(func(any types.Any) types.Any {
    ms := any.([]monad.Mono)
    var array []*message.Product
    for _, m := range ms {
      resp, e := m.(monad.Mono).Unwrap()
      if e != nil {
        logrus.Errorf("seach error: %s", e)
      } else {
        prod := resp.(*entity.Product)
        array = append(array, MapProductToMessage(prod))
      }
    }
    return array
  }).Then(func(any types.Any) types.Any {
    return &message.Products{Products: any.([]*message.Product)}
  })
}

