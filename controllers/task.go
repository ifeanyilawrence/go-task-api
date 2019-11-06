package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ifeanyilawrence/go-task-api/models"
	"github.com/ifeanyilawrence/go-task-api/repository"
	"gopkg.in/mgo.v2/bson"
)

//CreateTask : creats a New Task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	newTask, err := repository.CreateTask(task)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

//GetTasks : returns all the tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.AllTasks()

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(tasks)
}

//GetSingleTask : returns one task
func GetSingleTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	taskID := bson.ObjectIdHex(params["id"])

	task, err := repository.GetSingleTask(taskID)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(task)
}

//UpdateTask : modify task details
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	task.ID = bson.ObjectIdHex(params["id"])

	updatedTask, err := repository.UpdateTask(task)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(updatedTask)
}

//DeleteTask : removes a task from DB
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	taskID := bson.ObjectIdHex(params["id"])

	deleted, err := repository.DeleteTask(taskID)

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
