package db

import (
  "fmt"
  "github.com/go-pg/pg/v10"
  "github.com/google/uuid"
  "github.com/mlambda-net/monads/monad"
  "github.com/mlambda-net/net/pkg/spec"
  "github.com/mlambda-net/store/pkg/domain/entity"
  "github.com/mlambda-net/store/pkg/domain/repository"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "os"
  "strconv"
)

type storeRepository struct {
  db    *pg.DB
  cache StoreCache
  index StoreIndex
}

func (i *storeRepository) DeleteImages(id uuid.UUID) monad.Mono {
  _, err := i.db.Exec("call delete_images(?)", id)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(id)
}

func (i *storeRepository) Delete(id uuid.UUID) monad.Mono {
  _, err := i.db.Exec("call delete_product(?)", id)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(id)
}

func (i *storeRepository) Save(product *entity.Product) monad.Mono {
  _, err := i.db.Exec("call edit_product(?, ?, ?)", product.ID, product.Name, product.Description)
  if err != nil {
    return monad.ToMono(err)
  }
  i.index.Index(product.ID.String(), product)
  return monad.ToMono(product)
}

func (i *storeRepository) AddPreview(product *entity.Product) monad.Mono {
  product.Thumbnail.Id = uuid.New()
  _, err := i.db.Exec("call add_product_preview(?, ?, ?)", product.ID, product.Thumbnail.Id, product.Thumbnail.Path)
  if err != nil {
    return monad.ToMono(err)
  }
  return monad.ToMono(product)
}

func (i *storeRepository) AddImages(product *entity.Product) monad.Mono {
  var errors []error
  for _,img := range product.Images {
    img.Id = uuid.New()
    _, err := i.db.Exec("call add_product_image(?, ?, ?)", product.ID, img.Id, img.Path)
    if err != nil {
      errors = append(errors, err)
    }
  }
  if errors != nil {
    return monad.ToMono(fmt.Errorf("%s", errors))
  }

  return monad.ToMono(product)
}

func (i *storeRepository) Add(product *entity.Product) monad.Mono {
  _, err := i.db.Exec("call add_product(?, ?, ?, ?, ?, ?, ?, ?)",
    product.ID,
    product.Code,
    product.Name,
    product.Currency,
    product.Brand,
    product.Category,
    product.Description,
    product.Price)
  if err != nil {
    return monad.ToMono(err)
  }
  i.index.Index(product.ID.String(), product)
  return monad.ToMono(product)
}

func (i *storeRepository) Get(id uuid.UUID) monad.Maybe {
  prod, err := i.get(id)
  if err != nil {
    return monad.Unit(monad.ToMono(err))
  }

  err = i.loadImages(prod)
  if err != nil {
    return monad.Unit(monad.ToMono(err))
  }
  if prod != nil {
    return monad.Unit(monad.ToMono(prod))
  }

  return monad.Empty()
}

func (i *storeRepository) GetAll() []monad.Maybe {
  var monads []monad.Maybe
  var items []*entity.Product
  _, err := i.db.Query(&items, "SELECT id, code, name, description, currency, category, brand, price FROM product")
  if err != nil {
    monads = append(monads, monad.Unit( monad.ToMono(err)))
    return monads
  }

  for _, item := range items {
    err = i.loadImages(item)
    if err == nil {
      monads = append(monads, monad.Unit(monad.ToMono(item)))
    }
  }

  return monads
}

func (i *storeRepository) Search(value string) []monad.Maybe {
  ids := i.index.Search(value)
  monads := make([]monad.Maybe,0)
  for _, id := range ids {
    monads = append(monads, i.Get(id))
  }

  return monads
}

func (i *storeRepository) Query(spec spec.Spec) monad.Maybe {
  var items []entity.Product
  _, err := i.db.Query(&items,spec.Query("SELECT id, code, name, description, currency, category, brand, price  FROM product"), spec.Data()...)
  if err != nil {
    return monad.Unit(monad.ToMono(err))
  }

  if len(items) > 0 {
    return monad.Unit(monad.ToMono(&items[0]))
  }

  return monad.Empty()
}

func (i *storeRepository) Close() {
  _ = i.db.Close()
}

func (i *storeRepository) get(id uuid.UUID) (*entity.Product, error)   {
  var items []entity.Product
  query := fmt.Sprintf(`SELECT id, code, name, description, price FROM product Where id = '%s'`,  id.String())
  _, err := i.db.Query(&items, query)
  if err != nil {
    return nil, err
  }

  if len(items) > 0 {
    return &items[0], nil
  }

  return nil, nil
}

func (i *storeRepository) fetchPreview(id uuid.UUID) (*entity.Image, error)  {
  key := fmt.Sprintf("preview-%s", id)
  var img *entity.Image
  i.cache.Get(key, &img)
  if img != nil {
    return img, nil
  }

  var imgs []entity.Image
  _, err := i.db.Query(&imgs,fmt.Sprintf("SELECT image.* FROM image INNER JOIN product p on image.id = p.thumbnail WHERE p.id = '%s'", id))
  if err != nil {
    return  nil, err
  }
  if len(imgs) > 0 {
    im := &imgs[0]
    i.cache.Set(key, im, 0)
    return im, nil
  }
  return nil, nil
}

func (i *storeRepository) fetchImages(id uuid.UUID) ([]*entity.Image, error) {
  var images []*entity.Image
  _, err := i.db.Query(&images, fmt.Sprintf("SELECT image.* FROM image INNER JOIN product_image pi ON image.id = pi.image_id WHERE product_id = '%s'", id))

  if err != nil {
    return nil, err
  }

  return images, nil

}

func (i *storeRepository) loadImages(prod *entity.Product) error {
  if prod != nil {
    prev, err := i.fetchPreview(prod.ID)
    if err != nil {
      return err
    }

    images, err := i.fetchImages(prod.ID)
    if err != nil {
      return err
    }
    prod.Thumbnail = prev
    prod.Images = images
  }
  return nil
}

func NewStoreRepository(config *utils.Configuration) repository.StoreRepository {
  debug, _ := strconv.ParseBool( os.Getenv("Debug"))
  if debug {
    return &storeMock{}
  }
  db := initializeDB(config)
  cache := NewCache(config)
  index := NewIndex(config)

  return  &storeRepository{db: db, cache: cache, index: index}

}


