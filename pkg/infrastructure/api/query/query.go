package query

import (
  "github.com/graphql-go/handler"
  "github.com/mlambda-net/net/pkg/net"
)

type Query interface {
  Register(r net.Route)
}

type control struct {
  product   net.Request
}

func (c *control) Register(r net.Route) {
  h := handler.New(&handler.Config{
    Schema: c.Schema(),
    Pretty: true,
    GraphiQL: true,
    Playground: true,
  })

  r.AddRoute("graphql", "/store/graphql",nil, false, "POST", h.ServeHTTP)
}

func NewQuery( product net.Request) Query  {
  return &control{product: product}
}
