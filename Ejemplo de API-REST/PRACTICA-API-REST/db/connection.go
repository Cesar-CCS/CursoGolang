package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DSN = "host=localhost user=postgres password=C3m0nchi2828 dbname=prueba port=5432"
var DSN = "host=localhost user=postgres password=C3m0nchi28 dbname=practica port=5432"

var DB *gorm.DB

func DBConection() {
	var error error

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB conected")
	}
}
