package models

import "github.com/google/uuid"

type Doctor struct {
	Id          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Departments []*Department `json:"department"`
}
