package repository

import (
	"errors"

	"github.com/ifeanyilawrence/go-task-api/config"
	"github.com/ifeanyilawrence/go-task-api/models"
	"gopkg.in/mgo.v2/bson"
)

//Create : create a new task
func Create(task models.Task) (models.Task, error) {
	if (models.Task{}) == task {
		return models.Task{}, errors.New("400. Bad Request")
	}

	tsk := models.Task{}
	_ = config.Tasks.Find(bson.M{"description": task.Description}).One(&tsk)

	if tsk.Description != "" {
		return models.Task{}, errors.New("400. task with this description already exist")
	}

	task.ID = bson.NewObjectId()

	config.Tasks.Insert(task)

	return task, nil
}

//AllTasks : returns all existing tasks
func AllTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	err := config.Tasks.Find(bson.M{}).All(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//GetSingleTask : returns one tasks
func GetSingleTask(taskID bson.ObjectId) (models.Task, error) {
	if !taskID.Valid() {
		return models.Task{}, errors.New("400. Bad Request")
	}
	task := models.Task{}
	err := config.Tasks.Find(bson.M{"_id": taskID}).One(&task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

//Updatetask : updates task details
func Updatetask(task models.Task) (models.Task, error) {
	if (models.Task{}) == task || !task.ID.Valid() {
		return models.Task{}, errors.New("400. Bad Request")
	}

	err := config.Tasks.Update(bson.M{"_id": task.ID}, &task)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

//Deletetask : removes a task from the DB
func Deletetask(taskID bson.ObjectId) (bool, error) {
	if !taskID.Valid() {
		return false, errors.New("400. Bad Request")
	}

	err := config.Tasks.Remove(bson.M{"_id": taskID})
	if err != nil {
		return false, err
	}
	return true, nil
}
