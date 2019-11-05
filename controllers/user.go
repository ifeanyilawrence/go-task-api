package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/models"
	"github.com/ifeanyilawrence/go-task-api/repository"
	"gopkg.in/mgo.v2/bson"
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

//GetSingleUser : returns one user
func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	userID := bson.ObjectIdHex(params["id"])

	user, err := repository.GetSingleUser(userID)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

//UpdateUser : modify user details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	user.ID = bson.ObjectIdHex(params["id"])

	updatedUser, err := repository.UpdateUser(user)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(updatedUser)
}

//DeleteUser : removes a user from DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	userID := bson.ObjectIdHex(params["id"])

	deleted, err := repository.DeleteUser(userID)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, http.StatusText(500), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
