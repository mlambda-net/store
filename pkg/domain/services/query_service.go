package services

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/specs"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/mlambda-net/store/pkg/infrastructure/db"
)

type QueryService interface {
  SearchById(filter string) monad.Maybe
  SearchByName( filter string) monad.Maybe
  SearchByCode( filter string) monad.Maybe
  Search(value string) []monad.Maybe
  Get(id uuid.UUID) monad.Maybe
  GetAll() []monad.Maybe
}

type queryService struct {
  repo repository.StoreRepository
}

func (s *queryService) GetAll() []monad.Maybe {
   return s.repo.GetAll()
}

func (s *queryService) Get(id uuid.UUID) monad.Maybe {
  return s.repo.Get(id)
}

func (s *queryService) SearchById(filter string) monad.Maybe {
  return s.repo.Query(spec.Specify(specs.ById(filter)))
}

func (s *queryService) SearchByName( filter string) monad.Maybe {
  return s.repo.Query(spec.Specify(specs.ByName(filter)))
}

func (s *queryService) SearchByCode( filter string) monad.Maybe {
 return s.repo.Query(spec.Specify(specs.ByCode(filter)))
}

func (s *queryService) Search(value string) []monad.Maybe {
  return s.repo.Search(value)
}

func NewQueryService(config *utils.Configuration) QueryService {
  repo := db.NewStoreRepository(config)
  return &queryService{
    repo:  repo,
  }
}
