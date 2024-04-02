package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Cesar-CCS/PRACTICA-API-REST/db"
	"github.com/Cesar-CCS/PRACTICA-API-REST/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	json.NewDecoder(r.Body).Decode(&users)

	findedUsers := db.DB.Find(&users)
	err := findedUsers.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	//w.Write([]byte("GetUsersHandler"))

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// Para poder procesar informacion, necesitamos recibir un dato a traves de un objeto json
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	//leer request body
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DelteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User dont exist"))
		return
	}

	db.DB.Delete(&user, params["id"])

	json.NewEncoder(w).Encode(&user)
	w.WriteHeader(http.StatusOK)
}
