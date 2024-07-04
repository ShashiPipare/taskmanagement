package user

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id",json:"_id"`
	FirstName      string             `bson:"first_name",json:"first_name"`
	LastName       string             `bson:"last_name",json:"last_name"`
	Email          string             `bson:"email_id",json:"email_id"`
	Country        string             `bson:"country",json:"country"`
	Authentication Auth               `bson:"authentication",json:"authentication"`
}

type Auth struct {
	Token string `bson:"token",json:"token"`
	Hash  string `bson:"hash",json:"hash"`
}

type Claims struct {
	userID string `json:"_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
