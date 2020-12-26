package services

import (
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/specs"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/mlambda-net/store/pkg/infrastructure/db"
)

type LookupQuery interface {
  Categories() monad.Mono
  Brands() monad.Mono
}

type lookupQuery struct {
  repo repository.LookupRepository
}

func (l lookupQuery) Brands() monad.Mono {
  return l.repo.GetAll(spec.Specify(specs.ByLookupParent("Brand")), "Brand")
}

func (l lookupQuery) Categories() monad.Mono {
  return l.repo.GetAll(spec.Specify(specs.ByLookupParent("Category")), "Category")
}

func NewLookupQuery(config *utils.Configuration)  LookupQuery {
  repo := db.NewLookupRepository(config)
  return &lookupQuery{
    repo:  repo,
  }
}
