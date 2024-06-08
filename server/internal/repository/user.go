package repository

import "bocchikitsunei_webportfolio/internal/entities"

type UserRepository interface {
	GetAllUser() ([]entities.User, error)
	GetUserById(int) (*entities.User, error)
	GetUserByParams(int) (*entities.User, error)

	////////////////////////////////////////////////////////////////////

	CreateUser(user *entities.User) error
	GetUserByUsername(username string) (*entities.User, error)
}
