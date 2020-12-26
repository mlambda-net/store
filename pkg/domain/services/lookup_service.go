package services

import (
  "fmt"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/ex"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/specs"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/mlambda-net/store/pkg/infrastructure/db"
)

type LookupService interface {
  Add(item *entity.Lookup) monad.Mono
  Delete(name string) monad.Mono
	AddCategory(name string) monad.Mono
  AddBrand(name string) monad.Mono
  SubCategories(category string) monad.Maybe
}

type lookupService struct {
	repo  repository.LookupRepository
}

func (s *lookupService) SubCategories(category string) monad.Maybe {
  return s.repo.Get(spec.Specify(specs.ByLookupName(category))).Bind(func(any types.Any) monad.Maybe {
    c := any.(*entity.Lookup)
    if c.Parent == "Category" {
      return monad.Unit( s.repo.GetAll(spec.Specify(specs.ByLookupParent(category)), category))
    }
     return monad.Unit(monad.ToMono(ex.ToError( fmt.Sprintf("%s is not a kind of category", category))))
  })
}

func (s *lookupService) AddBrand(name string) monad.Mono {
  return s.repo.Add(&entity.Lookup{
    Name:   name,
    Parent: "Brand",
  })
}

func (s *lookupService) AddCategory(name string) monad.Mono {
  return s.repo.Add(&entity.Lookup{
    Name:   name,
    Parent: "Category",
  })
}

func (s *lookupService) Delete(name string) monad.Mono {
  return s.repo.Delete(name)
}

func (s *lookupService) Add(item *entity.Lookup) monad.Mono {
   return s.repo.Add(item)
}

func NewLookupService(config *utils.Configuration) LookupService {
  repo := db.NewLookupRepository(config)
  return &lookupService{
    repo:  repo,
  }
}
