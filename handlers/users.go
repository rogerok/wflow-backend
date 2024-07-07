package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/models"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
		Id:        "12sss3",
		LastName:  "1234",
	}

	return c.JSON(user)

}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Jhon doe")

}
