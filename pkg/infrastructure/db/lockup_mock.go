package db

import (
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

type brandMock struct {
}

func (b brandMock) TryAdd(lockup string, parent string) monad.Mono {
  panic("implement me")
}

func (b brandMock) Get(specify spec.Spec) monad.Maybe {
  return monad.Unit(&entity.Lookup{})
}

func (b brandMock) Add(product *entity.Lookup) monad.Mono {
  return monad.ToMono(product)
}

func (b brandMock) Delete(name string) monad.Mono {
  return monad.ToMono(name)
}

func (b brandMock) GetAll(specify spec.Spec, cacheKey string) monad.Mono {
  return monad.ToMono(make([]*entity.Lookup, 0))
}

func (b brandMock) Close() {
}
