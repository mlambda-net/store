package db

import (
  "encoding/json"
  "fmt"
  "github.com/go-pg/pg/v10"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/sirupsen/logrus"
)

func initializeDB(config *utils.Configuration) *pg.DB {
  db := pg.Connect(&pg.Options{
    User:     config.Db.User,
    Password: config.Db.Password,
    Addr:     fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
    Database: config.Db.Schema,
  })
  return db
}

func  serialize(obj interface{}) []byte  {
  b, err := json.Marshal(obj)
  if err != nil {
    logrus.Error(err)
    return nil
  }
  return b
}

func deserialize()  {
  
}
