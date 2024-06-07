package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
)

type UserService interface {
	GetUsers() ([]entities.User, error)
	GetUserId(int) (*entities.User, error)
	GetUserParams(int) (*entities.User, error)

	////////////////////////////////////////////////////////////////////

	GetEditUserProfile(int) (*entities.User, error)
	UpdateEditUserProfile(int, dtos.EditUserProfileRequest) (*entities.User, error)

	Register(request dtos.RegisterRequest) (*dtos.UserInfoResponse, error)
	Login(request dtos.LoginRequest, jwtSecret string) (*dtos.LoginResponse, error)
}
