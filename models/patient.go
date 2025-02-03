package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/spyder7370/SMH-Service/connections"
	"github.com/spyder7370/SMH-Service/encrypt"
)

type Patient struct {
	Id       uuid.UUID `json:"Id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Password string    `json:"-"`
	DOB      time.Time `json:"dob"`
	Phone    string
	Email    string
	Address  string
}

func (p *Patient) BookAppointment(doctor *Doctor, startTime time.Time, endTime time.Time, notes string) error {
	hashedPassword, err := encrypt.Hash(p.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO PATIENTS (name,password,dob,phone,email,address) values ($1,$2,$3,$4,$5,$6)`
	connections.DB.Exec(query, p.Name, hashedPassword)
	return nil
}
