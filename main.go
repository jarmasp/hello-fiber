package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Id        string
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "Jane",
		Lastname:  "Doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.NewString()
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	app.Get("/user", handleUser)
	app.Post("/user", handleCreateUser)
	app.Listen(":3000")
}
