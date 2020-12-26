package query

import (
  "github.com/graphql-go/graphql"
  "github.com/mlambda-net/store/pkg/application/message"
)

func (c *control) ProductByCodeArgs() graphql.FieldConfigArgument   {
  return graphql.FieldConfigArgument{
    "text": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  }
}

func (c *control) ProductByNameArgs() graphql.FieldConfigArgument   {
  return graphql.FieldConfigArgument{
    "name": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  }
}

func (c *control) ProductByTextArgs() graphql.FieldConfigArgument   {
  return graphql.FieldConfigArgument{
    "text": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  }
}

func (c *control) ProductType() *graphql.Object   {
  return graphql.NewObject(graphql.ObjectConfig{
    Name: "product",
    Fields: graphql.Fields{
      "id" : &graphql.Field{
        Type:              graphql.String,
      },
      "code" : &graphql.Field{
        Type: graphql.String,
      },
      "name" : &graphql.Field{
        Type: graphql.String,
      },
      "description" : &graphql.Field{
        Type: graphql.String,
      },
      "price" : &graphql.Field{
        Type: graphql.Float,
      },
      "preview" : &graphql.Field{
        Type: graphql.String,
      },
      "images" : &graphql.Field{
        Type: graphql.NewList(graphql.String),
      },
    },
  })
}

func (c *control) ProductResolveByCode ()   func (params graphql.ResolveParams) (interface{}, error) {
  return func(params graphql.ResolveParams) (interface{}, error) {
    code, ok := params.Args["code"].(string)
    if ok {
      res, err := c.product.Request(&message.Fetch{
        Filter: code,
        By:     message.CODE,
      }).Unwrap()

      if err != nil {
        return nil, err
      }
      return res.(*message.Products).Products, nil
    }
    return message.Product{}, nil
  }
}

func (c *control) ProductResolveByName ()   func (params graphql.ResolveParams) (interface{}, error) {
  return func(params graphql.ResolveParams) (interface{}, error) {
    code, ok := params.Args["code"].(string)
    if ok {
      res, err := c.product.Request(&message.Fetch{
        Filter: code,
        By:     message.NAME,
      }).Unwrap()

      if err != nil {
        return nil, err
      }
      return res.(*message.Products).Products, nil
    }
    return message.Product{}, nil
  }
}

func (c *control) ProductResolveByText ()   func (params graphql.ResolveParams) (interface{}, error) {
  return func(params graphql.ResolveParams) (interface{}, error) {
    text, ok := params.Args["text"].(string)
    if ok {
      res, err := c.product.Request(&message.Search{Value: text}).Unwrap()

      if err != nil {
        return nil, err
      }
      return res.(*message.Products).Products, nil
    }
    return message.Product{}, nil
  }
}

