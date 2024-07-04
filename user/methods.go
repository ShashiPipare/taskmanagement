package user

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/pbkdf2"
	"main.go/connection"
)

func (u *User) getByEmail(email string) (err error) {
	filter := bson.D{
		{
			Key:   "email_id",
			Value: email,
		},
	}
	err = connection.MI.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return
	}
	return
}
func (u *User) insertOne() (err error) {
	var insertedResult *mongo.InsertOneResult
	insertedResult, err = connection.MI.DB.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		return
	}
	u.ID, _ = insertedResult.InsertedID.(primitive.ObjectID)
	return
}

func (u *User) encryptPassword(password string) (err error) {
	salt := ""
	salt, err = generateRandomString(32)
	if err != nil {
		return
	}

	hash := pbkdf2.Key([]byte(password), []byte(salt), 872791, 64, sha512.New)
	byteKey := []byte(fmt.Sprintf("%s", hash))
	u.Authentication.Hash = hex.EncodeToString(byteKey)
	u.Authentication.Token = salt
	return
}

func (u *User) authenticateUser(password string) (err error, flag bool) {
	err = u.getByEmail(u.Email)
	if err != nil {
		return
	}
	rawHash := pbkdf2.Key([]byte(password), []byte(u.Authentication.Token), 872791, 64, sha512.New)
	byteKey := []byte(fmt.Sprintf("%s", rawHash))
	hash := hex.EncodeToString(byteKey)
	if hash != u.Authentication.Hash {
		return
	}
	flag = true
	return
}

func (u *User) createToken() (string, error) {
	claims := Claims{}
	claims.userID = u.ID.Hex()
	claims.Email = u.Email
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token, err := generateJWT(claims)
	return token, err
}
