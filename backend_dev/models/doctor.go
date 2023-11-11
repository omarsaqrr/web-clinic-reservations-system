package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model

	Email       string `gorm:"unique"`
	Password    string
	Role        string
	DoctorSlots []DoctorSlots `gorm:"foreignKey:DoctorID"`
}
