package tasks

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main.go/data"
)

func add(c *fiber.Ctx) (err error) {
	a := data.New(c)
	task := Task{}
	err = c.BodyParser(&task)
	if err != nil {
		log.Println("error in parsing body:", err)
		return a.Error(err)
	}
	log.Println("task:", task)
	err = task.add()
	if err != nil {
		log.Println("error in inserting task into db:", err)
		return a.Error(err)
	}
	return a.Data(task)
}

func update(c *fiber.Ctx) (err error) {
	a := data.New(c)
	timeStamp := time.Now().UTC()
	task := Task{}
	err = c.BodyParser(&task)
	if err != nil {
		log.Println("error in parsing body:", err)
		return a.Error(err)
	}
	task.Updated.Time = timeStamp
	err = task.update()
	if err != nil {
		log.Println("error in updating task:", err)
		return a.Error(err)
	}
	return a.Data(task)
}

func getByID(c *fiber.Ctx) (err error) {
	a := data.New(c)
	task := Task{}
	task.ID, _ = primitive.ObjectIDFromHex(c.Params("id", ""))
	if err != nil {
		log.Println("error in parsing body:", err)
		return a.Error(err)
	}
	if task.ID == primitive.NilObjectID {
		log.Println("nil objectID passed")
		return a.Error(ErrNilObjectID)
	}
	err = task.getOne()
	if err != nil {
		log.Println("error in fetching a task:", err)
		return a.Error(err)
	}
	return a.Data(task)
}

func getAllTasks(c *fiber.Ctx) (err error) {
	a := data.New(c)
	tasks := []Task{}
	tasks, err = getAll()
	if err != nil {
		log.Println("error in fetching tasks:", err)
		return a.Error(err)
	}
	return a.Data(tasks)
}

func delete(c *fiber.Ctx) (err error) {
	a := data.New(c)
	task := Task{}
	task.ID, _ = primitive.ObjectIDFromHex(c.Params("id", ""))
	if err != nil {
		log.Println("error in parsing body:", err)
		return a.Error(err)
	}
	if task.ID == primitive.NilObjectID {
		log.Println("nil objectID passed")
		return a.Error(ErrNilObjectID)
	}
	err = task.delete()
	if err != nil {
		log.Println("error while deleting a task:", err)
		return a.Error(err)
	}
	return a.Success()
}
