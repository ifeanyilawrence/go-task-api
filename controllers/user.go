package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ifeanyilawrence/go-task-api/models"
	"github.com/ifeanyilawrence/go-task-api/repository"
)

//CreateUser : creats a New User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	newUser, err := repository.Create(user)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

//GetUsers : returns all the users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.AllUsers()

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}
