package repository

import (
	"errors"

	"github.com/ifeanyilawrence/go-task-api/config"
	"github.com/ifeanyilawrence/go-task-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

//CreateUser : create a new user
func CreateUser(user models.User) (models.User, error) {
	if (models.User{}) == user {
		return models.User{}, errors.New("400. Bad Request")
	}

	usr := models.User{}
	_ = config.Users.Find(bson.M{"email": user.Email}).One(&usr)

	if usr.Email != "" {
		return models.User{}, errors.New("User with this email already exist")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

	if err != nil {
		return models.User{}, errors.New("Error While Hashing Password, Try Again")
	}

	user.Password = string(hash)

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

//GetSingleUser : returns one users
func GetSingleUser(userID bson.ObjectId) (models.User, error) {
	if !userID.Valid() {
		return models.User{}, errors.New("400. Bad Request")
	}
	user := models.User{}
	err := config.Users.Find(bson.M{"_id": userID}).One(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//UpdateUser : updates user details
func UpdateUser(user models.User) (models.User, error) {
	if (models.User{}) == user || !user.ID.Valid() {
		return models.User{}, errors.New("400. Bad Request")
	}

	err := config.Users.Update(bson.M{"_id": user.ID}, &user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

//DeleteUser : removes a user from the DB
func DeleteUser(userID bson.ObjectId) (bool, error) {
	if !userID.Valid() {
		return false, errors.New("400. Bad Request")
	}

	err := config.Users.Remove(bson.M{"_id": userID})
	if err != nil {
		return false, err
	}
	return true, nil
}
