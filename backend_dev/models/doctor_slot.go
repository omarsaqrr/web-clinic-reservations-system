package models

import (
	"gorm.io/gorm"
)

type DoctorSlots struct {
	gorm.Model
	DoctorID    uint
	DoctorEmail string

	Date     string
	Hour     string
	ISBooked bool
}
