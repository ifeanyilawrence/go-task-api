package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//DB exported database instance
var DB *mgo.Database

//Users collection
var Users *mgo.Collection

//Tasks collection
var Tasks *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("go-task-api")
	Users = DB.C("Users")
	Tasks = DB.C("Tasks")

	fmt.Println("Connected to mongo database.")
}
