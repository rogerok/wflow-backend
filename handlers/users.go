package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/models"
	"github.com/rogerok/wflow-backend/stores"
)

type UserHandler struct {
	userStore stores.UserStore
}

func (h *UserHandler) UsersList(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "John",
		Id:        "12sss3",
		LastName:  "1234",
	}

	return c.JSON(user)

}

//func NewUserHandler(userStore stores.UserStore) *UserHandler {
//	return &UserHandler{
//		userStore: userStore,
//	}
//}

func (h *UserHandler) UserById(c *fiber.Ctx) error {

	id := c.Params("id")
	user, err := h.userStore.GetUserById(id)
	if err != nil {
		return err
	}

	fmt.Println(id)

	return c.JSON(user)

}
