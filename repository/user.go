package repository

import (
	"errors"

	"github.com/ifeanyilawrence/go-task-api/config"
	"github.com/ifeanyilawrence/go-task-api/models"
	"gopkg.in/mgo.v2/bson"
)

//Create : create a new user
func Create(user models.User) (models.User, error) {
	if (models.User{}) == user {
		return models.User{}, errors.New("400. Bad Request")
	}

	usr := models.User{}
	_ = config.Users.Find(bson.M{"email": user.Email}).One(&usr)

	if usr.Email != "" {
		return models.User{}, errors.New("400. User with this email already exist")
	}

	user.ID = bson.NewObjectId()

	config.Users.Insert(user)

	return user, nil
}

//AllUsers : returns all existing users
func AllUsers() ([]models.User, error) {
	users := []models.User{}
	err := config.Users.Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
