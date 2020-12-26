package db

import (
  "github.com/google/uuid"
  "github.com/mlambda-net/monads/monad"
	"github.com/mlambda-net/net/pkg/spec"
	"github.com/mlambda-net/store/pkg/domain/entity"
)

type storeMock struct {
}

func (t *storeMock) GetAll() []monad.Maybe {
  panic("implement me")
}

func (t *storeMock) DeleteImages(id uuid.UUID) monad.Mono {
	panic("implement me")
}

func (t *storeMock) Delete(id uuid.UUID) monad.Mono {
	panic("implement me")
}

func (t *storeMock) Query(specify spec.Spec) monad.Maybe {
  return monad.Unit(&entity.Product{})
}

func (t *storeMock) Search(value string) []monad.Maybe {
  panic("implement me")
}

func (t *storeMock) Save(product *entity.Product) monad.Mono {
  panic("implement me")
}

func (t *storeMock) FetchPreview(id uuid.UUID) monad.Mono {
	panic("implement me")
}

func (t *storeMock) FetchImages(id uuid.UUID) monad.Mono {
	panic("implement me")
}

func (t *storeMock) AddPreview(product *entity.Product) monad.Mono {
  return monad.ToMono(product)
}

func (t *storeMock) AddImages(product *entity.Product) monad.Mono {
  return monad.ToMono(product)
}

func (t *storeMock) Add(product *entity.Product) monad.Mono {
	return monad.ToMono(product)
}

func (t *storeMock) Single(_ spec.Spec) monad.Mono {
	return monad.ToMono(&entity.Product{})
}

func (t *storeMock)  Get(id uuid.UUID) monad.Maybe  {
	return monad.Empty()
}

func (t *storeMock) Close() {
}
