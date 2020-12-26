package services

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/mlambda-net/store/pkg/infrastructure/db"
)

type StoreService interface {
  EditProduct(product *entity.Product) monad.Maybe
  Create(product *entity.Product) monad.Mono
	Add(product *entity.Product) monad.Mono
  Delete(id uuid.UUID) monad.Mono
}

type storeService struct {
	repo  repository.StoreRepository
}

func (s *storeService) Delete(id uuid.UUID) monad.Mono {
  return s.repo.DeleteImages(id).Bind(
    func(_ types.Any) monad.Mono {
      return s.repo.Delete(id)
    })
}

func (s *storeService) Add(product *entity.Product) monad.Mono {
   return s.repo.Add(product)
}

func (s *storeService) EditProduct(product *entity.Product) monad.Maybe {
  return s.repo.Get(product.ID).Bind(func(any types.Any) monad.Maybe {
    resp, err := any.(monad.Mono).Unwrap()
    if err != nil {
      return monad.Unit(monad.ToMono(err))
    }
    p := resp.(*entity.Product)
    p.Name = product.Name
    p.Description = product.Description
    return monad.Unit(s.repo.Save(p))
  })
}

func (s *storeService) Create(product *entity.Product) monad.Mono {
  return s.repo.Add(product).Bind(s.AddPreview).Bind(s.AddImages)
}

func (s *storeService) AddPreview(any types.Any) monad.Mono {
  return s.repo.AddPreview(any.(*entity.Product))
}

func (s *storeService) AddImages(any types.Any) monad.Mono {
  return s.repo.AddImages(any.(*entity.Product))
}


func NewStoreService(config *utils.Configuration) StoreService {
  repo := db.NewStoreRepository(config)

  return &storeService{
    repo:  repo,
  }
}
