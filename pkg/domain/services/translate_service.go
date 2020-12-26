package services

import (
  "cloud.google.com/go/translate"
  "context"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/sirupsen/logrus"
  "golang.org/x/text/language"
  "google.golang.org/api/option"
)

type TranslateService interface {
  ToLang(name string, lang string) monad.Mono
}

type translateService struct {
  client *translate.Client
}

func (t *translateService) ToLang(name string, lang string) monad.Mono {

  if lang == "es" {
    if t.client != nil {
      ctx := context.Background()
      labels, err := t.client.Translate(ctx, []string{name}, language.Spanish, &translate.Options{
        Source: language.English,
        Format: translate.Text,
      })

      if err != nil {
        return monad.ToMono(err)
      }

      return monad.ToMono( labels[0].Text)
    }
  }
  return monad.ToMono(name)
}

func NewTranslate(config *utils.Configuration) TranslateService  {
  ctx := context.Background()
  client, err := translate.NewClient(ctx, option.WithAPIKey(config.Google.Api))
  if err != nil {
    logrus.Error(err)
  }
  return &translateService{
    client : client,
  }
}
