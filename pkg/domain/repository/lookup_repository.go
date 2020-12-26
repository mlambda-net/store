package repository

import (
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

type LookupRepository interface {
  Add(item *entity.Lookup) monad.Mono
  Get(specify spec.Spec) monad.Maybe
  Delete(name string) monad.Mono
  GetAll(specify spec.Spec, cacheKey string) monad.Mono
  TryAdd(lockup string, parent string) monad.Mono
  Close()

}
