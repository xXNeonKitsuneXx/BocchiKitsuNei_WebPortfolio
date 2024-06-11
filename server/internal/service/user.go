package service

import (
	"bocchikitsunei_webportfolio/internal/entities"
)

type UserService interface {
	GetUsers() ([]entities.User, error)
	GetUserParams(int) (*entities.User, error)

}
