package user

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/pbkdf2"
	"main.go/connection"
)

func (u *User) encryptPassword(password string) error {
	salt, err := generateRandomString(32)
	if err != nil {
		return err
	}

	hash := pbkdf2.Key([]byte(password), []byte(salt), 872791, 64, sha512.New)
	byteKey := []byte(fmt.Sprintf("%s", hash))
	u.Authentication.Hash = hex.EncodeToString(byteKey)
	u.Authentication.Token = salt
	return nil
}

func (u *User) authenticateUser(password string) (error, bool) {
	filter := bson.M{"email_id": u.Email}
	err := connection.MI.DB.Collection("users").FindOne(context.Background(), filter).Decode(&u)
	if err != nil {
		return err, false
	}
	rawHash := pbkdf2.Key([]byte(password), []byte(u.Authentication.Token), 872791, 64, sha512.New)
	byteKey := []byte(fmt.Sprintf("%s", rawHash))
	hash := hex.EncodeToString(byteKey)
	if hash == u.Authentication.Hash {
		return nil, true
	}
	return nil, false
}
