package db

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/AsynkronIT/protoactor-go/log"
  "github.com/go-redis/redis/v8"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/sirupsen/logrus"
  "time"
)

type StoreCache interface {
  Get(key string, data interface{})
  Set(key string, value interface{}, time time.Duration)
  Delete(key string)
}

type storeCache struct {
  db *redis.Client
}

func (s storeCache) Delete(key string) {
  ctx, err := context.WithTimeout(context.Background(), 1*time.Second)
  if err != nil {
    logrus.Error(err)
  }
  s.db.Del(ctx,key)
}

func (s storeCache) Set(key string, data interface{}, time time.Duration) {
  ctx := context.Background()
  buffer := serialize(data)
  err := s.db.Set(ctx, key, buffer, time).Err()
  if err != nil {
    logrus.Error(err)
  }
}

func (s storeCache) Get(key string, value interface{}) {
  ctx := context.Background()
  data, err := s.db.Get(ctx, key).Result()
  if err != nil {
    log.Error(err)
    return
  }
  _ = json.Unmarshal([]byte(data), value)
}

func NewCache(config *utils.Configuration) StoreCache  {

  rdb := redis.NewClient(&redis.Options{
    Addr: fmt.Sprintf("%s:%d", config.Cache.Server, config.Cache.Port),
    Password: config.Cache.Password,
    DB: config.Cache.DB,
  })

  return storeCache{ db : rdb}
}
