package user

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"main.go/connection"
)

// Signup
func signUp(c *fiber.Ctx) (err error) {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Unable to parse the input parameters.",
			"error":   err,
		})
	}
	user := User{}
	if apiBody.Path("email_id").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "mandatory field not sent - email_id.",
		})
	}
	user.Email = apiBody.Path("email_id").Data().(string)
	filter := bson.M{"email_id": user.Email}
	err = connection.MI.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "unable to check if user exists.",
			"error":   err,
		})
	}
	if user.ID != primitive.NilObjectID {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "user already exists.",
		})
	}
	if apiBody.Path("first_name").Data() != nil {
		user.FirstName = apiBody.Path("first_name").Data().(string)
	}
	if apiBody.Path("last_name").Data() != nil {
		user.LastName = apiBody.Path("last_name").Data().(string)
	}
	if apiBody.Path("country").Data() != nil {
		user.Country = apiBody.Path("country").Data().(string)
	}
	if apiBody.Path("password").Data() != nil {
		password := apiBody.Path("password").Data().(string)
		err := user.encryptPassword(password)
		if err != nil {
			return c.Status(200).JSON(&fiber.Map{
				"status":  false,
				"message": "Error in generating random string",
				"error":   err,
			})
		}
	} else {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Password not sent",
		})
	}
	user.ID = primitive.NewObjectID()
	res, err := connection.MI.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "unable to sign up the user.",
			"error":   err,
		})
	}
	auth := Auth{}
	user.Authentication = auth
	return c.Status(200).JSON(&fiber.Map{
		"success":      true,
		"message":      "user signed up successfully.",
		"user_details": res,
	})
}

// login without jwt
func login(c *fiber.Ctx) (err error) {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Unable to parse the input parameters.",
			"error":   err,
		})
	}
	if apiBody.Path("email_id").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "email_id not sent.",
		})
	}
	if apiBody.Path("password").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "password not sent.",
		})
	}
	user := User{}
	user.Email = apiBody.Path("email_id").Data().(string)
	password := apiBody.Path("password").Data().(string)

	err, ok := user.authenticateUser(password)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "unable to authenicate user.",
			"error":   err,
		})
	}
	if !ok {
		return c.Status(200).JSON(&fiber.Map{
			"success": false,
			"message": "Incorrect password.",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "user logged in succesfully.",
	})
}

func logout(c *fiber.Ctx) (err error) {

	return
}
