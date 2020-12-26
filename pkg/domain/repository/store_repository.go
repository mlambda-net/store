package repository

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

type StoreRepository interface {
	Add(product *entity.Product) monad.Mono
	AddPreview(product *entity.Product) monad.Mono
  AddImages(product *entity.Product) monad.Mono
  Save(product *entity.Product) monad.Mono
  DeleteImages(id uuid.UUID) monad.Mono
  Delete(id uuid.UUID) monad.Mono
  Get(id uuid.UUID) monad.Maybe
	Search(value string) []monad.Maybe
  GetAll()  []monad.Maybe
	Query(specify spec.Spec) monad.Maybe
  Close()

}
