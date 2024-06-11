package service

import (
	"bocchikitsunei_webportfolio/internal/entities"
	"bocchikitsunei_webportfolio/internal/repository"
	"github.com/gofiber/fiber/v2"
	"log"
)

type userService struct {
	userRepo  repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{
		userRepo:  userRepo,
	}
}

func (s userService) GetUsers() ([]entities.User, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponses := []entities.User{}
	for _, user := range users {
		userResponse := entities.User{
			UserID:   user.UserID,
			Username: user.Username,
			Password: user.Password,
		}
		userResponses = append(userResponses, userResponse)
	}
	return userResponses, nil
}

func (s userService) GetUserParams(userid int) (*entities.User, error) {
	user, err := s.userRepo.GetUserByParams(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user.UserID == nil && user.Username == nil && user.Password == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "user data is not found")
	}

	userResponse := entities.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
	return &userResponse, nil
}
