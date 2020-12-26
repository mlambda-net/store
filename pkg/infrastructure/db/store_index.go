package db

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/elastic/go-elasticsearch/v8"
  "github.com/elastic/go-elasticsearch/v8/esapi"
  "github.com/google/uuid"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/sirupsen/logrus"
  "strings"
)

type StoreIndex interface {
  Search(data string)  []uuid.UUID
  Index(key string, value interface{})
}


type storeIndex struct {
  client *elasticsearch.Client
}

func (s storeIndex) Search(data string) []uuid.UUID {

  query := fmt.Sprintf( `{"query":{"query_string" :{ "query": "%s"}}}`, data)
  res, err := s.client.Search(
    s.client.Search.WithContext(context.Background()),
    s.client.Search.WithIndex("store"),
    s.client.Search.WithBody(strings.NewReader(query)),
  )

  if err != nil {
    logrus.Error(err)
  }

  if res.IsError() {
    logrus.Error("Error: %s", res.String())
  }

  ids := s.getIds(res)

  return ids
}

func (s storeIndex) getIds(res *esapi.Response) []uuid.UUID {
  var r map[string]interface{}
  if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
    logrus.Errorf("Error parsing the response body: %s", err)
  }

  var ids []uuid.UUID
  for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
    id, err := uuid.Parse(hit.(map[string]interface{})["_id"].(string))
    if err == nil {
      ids = append(ids, id)
    }

  }
  return ids
}

func (s storeIndex) Index(key string, value interface{}) {
  _, err := s.client.Index(
    "store",
    strings.NewReader(string(serialize(value))),
    s.client.Index.WithDocumentID(key),
    s.client.Index.WithRefresh("true"))

  if err != nil {
    logrus.Errorf("error indexing: %s", err)
  }
}

func NewIndex(config *utils.Configuration)  StoreIndex {
  client := elasticClient(config)
  return storeIndex{client: client}
}

func elasticClient(config *utils.Configuration) *elasticsearch.Client {

  var conf elasticsearch.Config
  if config.Index.Authenticate {
    conf = elasticsearch.Config{
      Addresses: []string{
        config.Index.Server,
      },
      Username: config.Index.User,
      Password: config.Index.Password,
    }
  } else {
    conf = elasticsearch.Config{
      Addresses: []string{
        config.Index.Server,
      },
    }
  }

  client, err := elasticsearch.NewClient(conf)
  if err != nil {
    logrus.Error(err)
  }

  return client
}
