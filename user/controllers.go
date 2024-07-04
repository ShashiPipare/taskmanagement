package user

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/data"
)

// Signup
func signUp(c *fiber.Ctx) (err error) {
	a := data.New(c)
	type RequestParams struct {
		User
		Password string `json:"password"`
	}
	params := RequestParams{}
	err = c.BodyParser(&params)
	if err != nil {
		return a.Error(err)
	}
	if params.Email == "" || params.Password == "" {
		return a.Error(ErrMissingRequiredParams)
	}
	user := User{}
	err = user.getByEmail(params.Email)
	if err != nil {
		return a.Error(err)
	}
	if user.ID != primitive.NilObjectID {
		return a.Error(ErrUserAlreadyExists)
	}
	user.Email = params.Email
	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Country = params.Country
	err = user.encryptPassword(params.Password)
	if err != nil {
		return a.Error(err)
	}

	err = user.insertOne()
	if err != nil {
		return a.Error(err)
	}

	auth := Auth{}
	user.Authentication = auth
	return a.Data(user)
}

// login without jwt
func login(c *fiber.Ctx) (err error) {
	a := data.New(c)
	type RequestParams struct {
		User
		Password string `json:"password"`
	}
	params := RequestParams{}
	err = c.BodyParser(&params)
	if err != nil {
		return a.Error(err)
	}
	if params.Email == "" || params.Password == "" {
		return a.Error(ErrMissingRequiredParams)
	}
	user := User{}
	user.Email = params.Email
	err, ok := user.authenticateUser(params.Password)
	if err != nil {
		return a.Error(err)
	}
	if !ok {
		return a.Error(ErrIncorrectPassword)
	}
	jwtToken, err := user.createToken()
	if err != nil {
		return a.Error(ErrAuthUser)
	}
	return a.Data(jwtToken)
}

func logout(c *fiber.Ctx) (err error) {

	return
}
