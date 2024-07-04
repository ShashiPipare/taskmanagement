package tasks

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/connection"
)

func getAll() (tasks []Task, err error) {
	var cur *mongo.Cursor
	filter := bson.D{}
	cur, err = connection.MI.DB.Collection(CollectionName).Find(context.TODO(), filter)
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), &tasks)
	if err != nil {
		return
	}
	return
}
