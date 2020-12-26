package services

import (
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/utils"
)

type CurrencyService interface {
  GetCurrencies() monad.Mono

}

type currencyService struct {

}

func (c currencyService) GetCurrencies() monad.Mono {
  return monad.ToMono([]*entity.Currency{
    &entity.Currency {
      Character: "$",
      Symbol: "USD",
      Name:   "Dollars",
    },
    &entity.Currency {
      Character: "₡",
      Symbol: "CRC",
      Name:   "Colones",
    },
  })
}

func NewCurrencyService(*utils.Configuration) CurrencyService  {
  return currencyService{}
}
