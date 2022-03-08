package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// UserList godoc
// @Summary Get all users
// @Description Get user list. Auth not required
// @ID get-user-lists
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} userDataListResponse
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /users/list [get]
func (h *Handler) UserList(c *fiber.Ctx) error {
	users, err := h.userStore.List()
	if err != nil {
		status := http.StatusInternalServerError
		message := fmt.Sprintf("[Error] retreiving article info: %v", err.Error())
		return c.Status(status).JSON(map[string]interface{}{"message": message})
	}

	status := c.Response().StatusCode()
	message := "success"
	return c.Status(status).JSON(newUserListResponse(h.userStore, users, status, message))
}
