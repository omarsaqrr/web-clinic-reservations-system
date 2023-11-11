package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID uint
	//SlotID       uint
	DoctorEmail  string
	PatientEmail string
	Date         string
	Hour         string
}
