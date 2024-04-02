package main

import (
	"net/http"

	"github.com/gorilla/mux"

	//Cada archivo es un paquete, asi importamos paqueteslocales
	"github.com/Cesar-CCS/PRACTICA-API-REST/db"
	"github.com/Cesar-CCS/PRACTICA-API-REST/models"
	"github.com/Cesar-CCS/PRACTICA-API-REST/routes"
)

func main() {

	db.DBConection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//r nos ayudara a agrupar multiples rutas
	//Este es el servidor
	r := mux.NewRouter()

	//Funcion de la ruta
	r.HandleFunc("/", routes.HomeHandler)

	//rutas de mis peticiones
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DelteUserHandler).Methods("DELETE")

	//Inicializando el servidor
	http.ListenAndServe(":3000", r)
}
