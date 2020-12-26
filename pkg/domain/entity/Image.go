package entity

import "github.com/google/uuid"

type Image struct {
  Id   uuid.UUID
  Path string
}
