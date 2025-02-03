package models

import "github.com/google/uuid"

type Department struct {
	Id   uuid.UUID
	Name string
}
