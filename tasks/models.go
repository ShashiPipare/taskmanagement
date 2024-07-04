package tasks

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var CollectionName string = "tasks"

type Task struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Created     Created            `json:"created" bson:"created"`
	Updated     Updated            `json:"updated" bson:"updated"`
	Deleted     Deleted            `json:"deleted" bson:"deleted"`
	Assignee    Assignee           `json:"assignee" bson:"assignee"`
}

type Created struct {
	Name string    `json:"name" bson:"name"`
	Time time.Time `json:"time" bson:"time"`
}
type Updated struct {
	Name string    `json:"name" bson:"name"`
	Time time.Time `json:"time" bson:"time"`
}
type Deleted struct {
	Ok   bool      `json:"ok" bson:"ok"`
	Name string    `json:"name" bson:"name"`
	Time time.Time `json:"time" bson:"time"`
}
type Assignee struct {
	Name   string             `json:"name" bson:"name"`
	UserID primitive.ObjectID `jspn:"user_id" bson:"user_id"`
}
