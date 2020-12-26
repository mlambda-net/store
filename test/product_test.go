package test

import (
  "github.com/mlambda-net/store/pkg/application/message"
  "github.com/mlambda-net/store/pkg/domain/utils"
  "github.com/stretchr/testify/assert"
  "testing"
)


func Test_Add_Product(t *testing.T) {
  msg := CreateProduct()
  c := NewClient()
  actor := c.LoadActor("store")
  r, err := actor.Request(msg).Unwrap()

  assert.Nil(t, err)
  assert.NotNil(t, r)
}

func Test_Fetch_By_Code_Product (t *testing.T)  {

  msg := FetchProduct("teverde100g", message.CODE)

  c := NewClient()
  actor := c.LoadActor("store")
  r, err := actor.Request(msg).Unwrap()

  assert.Nil(t, err)
  assert.NotNil(t, r)
}


func Test_Fetch_By_Name_Product (t *testing.T)  {
  msg := FetchProduct("Té", message.NAME)
  c := NewClient()
  actor := c.LoadActor("store")
  r, err := actor.Request(msg).Unwrap()

  assert.Nil(t, err)
  assert.NotNil(t, r)
}

func FetchProduct(value string, by message.SearchBy) *message.Fetch {
  return &message.Fetch{
    Filter: value,
    By:     by,
  }
}

func Test_Edit_Product(t *testing.T)  {
  msg := FetchProduct("teverde100g", message.CODE)
  c := NewClient()
  actor := c.LoadActor("store")
  r, _ := actor.Request(msg).Unwrap()
  p := r.(*message.Product)
  ed := EditProduct(p)
  _,  err:= actor.Request(ed).Unwrap()
  assert.NoError(t, err)
}

func Test_Search_Product(t *testing.T) {
  msg := &message.Search{Value: "Camellia"}
  c := NewClient()
  actor := c.LoadActor("store")
  r, _ := actor.Request(msg).Unwrap()
  assert.NotNil(t, r)
}



func Test_Edit_Images(t *testing.T)  {

}

func Test_Change_Price(t *testing.T)  {

}

func EditProduct(p *message.Product) *message.EditProduct {
  return &message.EditProduct{
    Id:          p.Id,
    Name:        "Change Name",
    Description: p.Description,
  }
}

func CreateProduct() *message.CreateProduct {
  return &message.CreateProduct{
    Code: "teverde100g",
    Name:        "Té Verde",
    Description: "El té verde es un producto derivado de la Camellia sinensis. Se usa como infusión o directamente como una medicina natural. Al té verde se le dan propiedades para mejorar la agilidad mental y el pensamiento, para perder peso y para el tratamiento de trastornos intestinales, los dolores de cabeza, la osteoporosis, etc…",
    Price:       5,
    Thumbnail: &message.Image{
      Image: utils.ReadBytesFromImage("./assets/teverde.jpeg"),
    },
    Images: []*message.Image{
      &message.Image{
        Image: utils.ReadBytesFromImage("./assets/teverde.jpeg"),
      },
      &message.Image{
        Image: utils.ReadBytesFromImage("./assets/te-verde-m.jpg"),
      },
    },
  }
}
