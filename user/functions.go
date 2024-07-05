package user

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/config"
)

var JWTKey string

func Init(conf config.Conf) {
	JWTKey = conf.JWTKey
}

/*
This function returns a URL-safe, base64 encoded securely generated random string
*/
func generateRandomString(size int) (string, error) {
	byteData, err := generateRandomBytes(size)
	return base64.URLEncoding.EncodeToString(byteData), err
}

/*
This function is used to securely generated random bytes
*/

func generateRandomBytes(size int) ([]byte, error) {
	token := make([]byte, size)
	_, err := rand.Read(token)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return token, nil
}

/*
This function is used to generate JWT token
*/
func generateJWT(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
This function validates JWT token
*/
func verifyToken(ID primitive.ObjectID, tokenString string) error {
	user := User{}
	err := user.getByID(ID)
	if err != nil {
		return err
	}
	if !user.IsActive {
		return fmt.Errorf(ErrInactiveUser.Error())
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf(ErrInvalidJWT.Error())
	}
	return nil
}
