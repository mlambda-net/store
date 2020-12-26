package services

import (
  "bytes"
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/mlambda-net/store/pkg/infrastructure/db"
)

type ImageService interface {
  AddImage(file *entity.File) monad.Mono
}

type imageService struct {
  loader *s3manager.Uploader
  repo   repository.StoreRepository
  bucket string
}

func (i *imageService) AddImage(file *entity.File)  monad.Mono {
  name := fmt.Sprintf("%s_%s", file.Prefix, file.Name)
  _, err := i.loader.Upload(&s3manager.UploadInput{
    Bucket: aws.String(i.bucket),
    Key:    aws.String(name),
    Body:   bytes.NewReader(file.Content),
  })
  if err != nil {
    return monad.ToMono(err)
  }

  return monad.ToMono(&entity.Image{
    Path: fmt.Sprintf("https://image-mlambda.sirv.com/%s", name),
  })
}

func NewImageService(config *utils.Configuration) ImageService {
  cred := credentials.NewStaticCredentials(config.S3.Key, config.S3.Secret, "")
  sess := session.Must(session.NewSession(&aws.Config{

    Credentials:                       cred,
    Endpoint:                          aws.String(config.S3.Endpoint),
    DisableSSL:                        aws.Bool(true),

  }))
  uploader := s3manager.NewUploader(sess)
  repo := db.NewStoreRepository(config)
  return &imageService{repo: repo, loader: uploader, bucket: config.S3.Bucket}
}
