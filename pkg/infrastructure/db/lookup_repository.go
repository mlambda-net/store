package db

import (
  "github.com/go-pg/pg/v10"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/specs"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "os"
  "strconv"
  "time"
)

type lookupRepository struct {
  db    *pg.DB
  cache StoreCache
}

func (r *lookupRepository) TryAdd(lockup string, parent string)  monad.Mono {
  mb := r.Get(spec.Specify(specs.ByLookupName(lockup)))
  switch m := mb.(type) {
  case monad.Just:
    return monad.ToMono(m.Value())
  default:
    item := &entity.Lookup{
      Name:   lockup,
      Parent: parent,
    }
    r.Add(item)
    return monad.ToMono(item)
  }
}

func (r *lookupRepository) Get(specify spec.Spec) monad.Maybe {
  items, err := r.lookups(specify)
  if err != nil {
    return monad.Unit(err)
  }
  if len(items) > 0 {
    return monad.Unit(items[0])
  }
  return monad.Empty()
}

func (r *lookupRepository) Add(item *entity.Lookup) monad.Mono {
  _, err := r.db.Exec("call lookup_add(?, ?)", item.Name, item.Parent)
  if err != nil {
    return monad.ToMono(err)
  }
  r.cache.Delete(item.Parent)
  return monad.ToMono(item)
}


func (r *lookupRepository) Delete(name string) monad.Mono {
  _, err := r.db.Exec("call lookup_delete(?)", name)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(name)
}

func (r *lookupRepository) GetAll(specify spec.Spec, cacheKey string) monad.Mono {
  var items []*entity.Lookup
  r.cache.Get(cacheKey, &items)
  if items != nil {
    return monad.ToMono(items)
  } else {
    values, err := r.lookups(specify)
    if err != nil {
      return monad.ToMono(err)
    }
    r.cache.Set(cacheKey, values, 1*time.Hour)
    return monad.ToMono(values)
  }
}

func (r *lookupRepository) lookups(specify spec.Spec) ([]*entity.Lookup, error) {
  var items []*entity.Lookup
  query := "SELECT name, parent FROM lookup"
  _, err := r.db.Query(&items, specify.Query(query), specify.Data()...)
  if err != nil {
    return nil, err
  }
  return items, err
}

func (r *lookupRepository) Close() {
  _ = r.db.Close()
}

func NewLookupRepository(config *utils.Configuration) repository.LookupRepository {
  debug, _ := strconv.ParseBool( os.Getenv("Debug"))
  if debug {
    return &brandMock{}
  }
  db := initializeDB(config)
  cache := NewCache(config)
  return  &lookupRepository{db: db, cache: cache}
}
