package models

import (
	"time"

	"github.com/google/uuid"
)

type AppointmentRequest struct {
	DoctorID  uuid.UUID `json:"doctor_id"`
	PatientId uuid.UUID `json:"patient_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type Appointment struct {
	StartTime time.Time
	EndTime   time.Time
	Patient   *Patient
	Doctor    *Doctor
}
