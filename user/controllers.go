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

func signUp(c *fiber.Ctx) (err error) {
	apiBody, err := gabs.ParseJSON([]byte(c.Body()))
	if err != nil {
		return err
	}
	user := User{}
	if apiBody.Path("phone_no").Data() == nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "mandatory field not sent - phone_no.",
		})
	}
	user.PhoneNo = apiBody.Path("phone_no").Data().(string)
	filter := bson.M{"phone_no": user.PhoneNo}
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
	user.ID = primitive.NewObjectID()
	res, err := connection.MI.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "unable to sign up the user.",
			"error":   err,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success":      true,
		"message":      "user signed up successfully.",
		"user_details": res,
	})
}

func login(c *fiber.Ctx) (err error) {

	return
}

func logout(c *fiber.Ctx) (err error) {

	return
}
