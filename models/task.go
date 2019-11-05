package models

import "gopkg.in/mgo.v2/bson"

//Task model
type Task struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Description string        `json:"description" bson:"description"`
	Completed   bool          `json:"completed" bson:"completed"`
}
