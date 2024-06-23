package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "John",
		Id:        "123",
		LastName:  "1234",
	}

	return c.JSON(user)

}
func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Jhon doe")

}
