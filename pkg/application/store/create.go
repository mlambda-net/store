package store

import (
  "github.com/google/uuid"
  types "github.com/mlambda-net/monads"
  "github.com/mlambda-net/net/pkg/core"
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func (a *storeActor) create(product *message.CreateProduct) core.Resolver {
  resolver := core.NewResolve()

  thumbnail := &entity.Image{
    Id:   uuid.New(),
    Path: product.Thumbnail,
  }

  images := make([]*entity.Image, 0)
  for _, i := range product.Images {
    images = append(images, &entity.Image{
      Id:   uuid.New(),
      Path: i,
    })
  }

  id, err := uuid.Parse(product.Id)

  if err != nil {
    return resolver.Error(err)
  }

  return resolver.Mono(a.service.Create(&entity.Product{
    ID:          id,
    Code:        product.Code,
    Name:        product.Name,
    Brand:       product.Brand,
    Category:    product.Category,
    Currency:    product.Currency,
    Description: product.Description,
    Price:       product.Price,
    Thumbnail:   thumbnail,
    Images:      images,
  })).Then(func(any types.Any) types.Any {
    return &core.Done{}
  })
}

