package models

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	Id       uuid.UUID `json:"Id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Password string    `json:"-"`
	DOB      time.Time `json:"dob"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
}
