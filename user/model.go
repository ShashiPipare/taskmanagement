package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id",json:"_id"`
	FirstName string             `bson:"first_name",json:"first_name"`
	LastName  string             `bson:"last_name",json:"last_name"`
	PhoneNo   string             `bson:"phone_no",json:"phone_no"`
	Country   string             `bson:"country",json:"country"`
}
