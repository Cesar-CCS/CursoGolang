package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	//Gorm nos permite especificar tipos y requerimientos especificos
	//Gorm tambien permite definir la estructura (json) que quiero que el cliente use
	FirstName string `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
