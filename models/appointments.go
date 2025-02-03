package models

import "time"

type Appointment struct {
	StartTime time.Time
	EndTime   time.Time
	Patient   *Patient
	Doctor    *Doctor
}
