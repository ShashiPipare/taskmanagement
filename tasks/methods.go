package tasks

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/connection"
)

func (t *Task) add() (err error) {
	var InsertedResponse *mongo.InsertOneResult
	InsertedResponse, err = connection.MI.DB.Collection(CollectionName).InsertOne(context.TODO(), t)
	if err != nil {
		log.Println("error in inserting a task:", err)
		return
	}
	t.ID, _ = InsertedResponse.InsertedID.(primitive.ObjectID)
	return
}

func (t *Task) update() (err error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: t.ID,
		},
	}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "name",
					Value: t.Name,
				},
				{
					Key:   "description",
					Value: t.Description,
				},
				{
					Key:   "updated",
					Value: t.Updated,
				},
				{
					Key:   "assignee",
					Value: t.Assignee,
				},
			},
		},
	}
	after := options.After
	upsert := false
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	err = connection.MI.DB.Collection(CollectionName).FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(t)
	if err != nil {
		return
	}
	return
}

func (t *Task) getOne() (err error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: t.ID,
		},
	}
	err = connection.MI.DB.Collection(CollectionName).FindOne(context.TODO(), filter).Decode(t)
	if err != nil {
		return
	}
	return
}

func (t *Task) delete() (err error) {
	filter := bson.D{
		{
			Key:   "_id",
			Value: t.ID,
		},
	}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "deleted",
					Value: t.Deleted,
				},
			},
		},
	}

	_, err = connection.MI.DB.Collection(CollectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}
