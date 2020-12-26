package query

import (
  "github.com/graphql-go/graphql"
)



func (c *control) SchemaType() *graphql.Object {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "Products",
    Fields: graphql.Fields{
      "product_by_code": &graphql.Field{
        Type:             graphql.NewList(c.ProductType()),
        Args:              c.ProductByCodeArgs(),
        Resolve:           c.ProductResolveByCode(),
      },
      "product_by_name": &graphql.Field{
        Type:             graphql.NewList( c.ProductType()),
        Args:              c.ProductByNameArgs(),
        Resolve:           c.ProductResolveByName(),
      },
      "search": &graphql.Field{
        Type:             graphql.NewList( c.ProductType()),
        Args:              c.ProductByTextArgs(),
        Resolve:           c.ProductResolveByText(),
      },
    },

  })
}

func (c control) Schema() *graphql.Schema  {
  var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query:        c.SchemaType(),
  })
  return &schema
}

