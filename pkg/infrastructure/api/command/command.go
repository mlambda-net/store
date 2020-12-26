package command

import (
  "github.com/mlambda-net/net/pkg/net"
)

type Command interface {
  Register(r net.Route)
}

type control struct {
  store    net.Request
  currency net.Request
  lookup   net.Request
  image   net.Request
}

func NewCommand(store, currency, lookup, image net.Request) Command {
  return &control{
    store: store,
    currency: currency,
    lookup: lookup,
    image: image,
  }
}

func (c *control) Register(r net.Route) {
  r.AddRoute("product_add", "/store/product", nil, false, "POST", c.CreateProduct)
  r.AddRoute("product_edit", "/store/product", nil, false, "PUT", c.EditProduct)
  r.AddRoute("product_delete", "/store/product/{id}", nil, false, "DELETE", c.DeleteProduct)
  r.AddRoute("product_get_stream", "/store/product/stream", nil, false, "GET", c.GetProductStream)
  r.AddRoute("product_get_all", "/store/product", nil, false, "GET", c.GetProductAll)
  r.AddRoute("product_get", "/store/product/{id}", nil, false, "GET", c.GetProduct)

  r.AddRoute("currency_get_all", "/store/currency", nil, false, "GET", c.CurrencyAll)
  r.AddRoute("category_get_all", "/store/category/{lang}", nil, false, "GET", c.CategoryAll)
  r.AddRoute("category_add", "/store/category", nil, false, "POST", c.CategoryAdd)
  r.AddRoute("brand_get_all", "/store/brand", nil, false, "GET", c.BrandAll)
  r.AddRoute("brand_add", "/store/brand", nil, false, "POST", c.BrandAdd)

  r.AddRoute("image_add", "/store/image", nil, false, "POST", c.ImageAdd)

}
