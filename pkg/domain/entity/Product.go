package entity

import (
  "fmt"
  "github.com/google/uuid"
)

type Product struct {
  ID          uuid.UUID
  Name        string
  Description string
  Price       float64
  Code        string
  Brand       string
  Category    string
  Currency    string
  Thumbnail   *Image
  Images      []*Image
}

func (p Product) GetPreview() string {
	if p.Thumbnail != nil {
	  return p.Thumbnail.Path
  }
  return ""
}

func (p Product) GetImages() []string {
  images := make([]string,0)
  for _, img := range p.Images {
    images = append(images, img.Path)
  }
  return images
}

func (p Product) PriceText() string {
  if p.Currency == "USD" {
    return fmt.Sprintf("$ %.2f", p.Price)
  }

  return fmt.Sprintf("₡ %.2f", p.Price)
}

