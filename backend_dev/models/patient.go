package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model

	Email       string `gorm:"unique"`
	Password    string
	Role        string
	Appointment []Appointment `gorm:"foreignKey:PatientID"`
}
