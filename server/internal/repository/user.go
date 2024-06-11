package repository

import "bocchikitsunei_webportfolio/internal/entities"

type UserRepository interface {
	GetAllUser() ([]entities.User, error)
	GetUserByParams(int) (*entities.User, error)

}
