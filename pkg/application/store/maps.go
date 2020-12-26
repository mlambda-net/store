package store

import (
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/entity"
)

func MapProductToMessage(prod *entity.Product) *message.Product {
  return &message.Product{
    Id:          prod.ID.String(),
    Code:        prod.Code,
    Name:        prod.Name,
    Description: prod.Description,
    Currency:    prod.Currency,
    Brand:       prod.Brand,
    Category:    prod.Category,
    Price:       prod.Price,
    Preview:     prod.GetPreview(),
    Images:      prod.GetImages(),
    PriceText:   prod.PriceText(),
  }
}


