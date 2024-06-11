package handler

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/service"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type userHandler struct {
	userSer   service.UserService
	uploadSer service.UploadService
}

func NewUserHandler(userSer service.UserService, uploadSer service.UploadService) userHandler {
	return userHandler{userSer: userSer, uploadSer: uploadSer}
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	usersResponse := make([]dtos.UsersResponse, 0)

	users, err := h.userSer.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		usersResponse = append(usersResponse, dtos.UsersResponse{
			UserID:   user.UserID,
			Username: user.Username,
			Password: user.Password,
		})
	}
	return c.JSON(usersResponse)
}

func (h *userHandler) GetUserParams(c *fiber.Ctx) error {
	userIDReceive, err := strconv.Atoi(c.Params("UserID"))

	user, err := h.userSer.GetUserParams(userIDReceive)
	if err != nil {
		return err
	}

	userResponse := dtos.UserParamsResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}

	return c.JSON(userResponse)
}
