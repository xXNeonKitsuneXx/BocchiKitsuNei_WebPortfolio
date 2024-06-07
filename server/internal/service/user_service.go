package service

import (
	"bocchikitsunei_webportfolio/internal/dtos"
	"bocchikitsunei_webportfolio/internal/entities"
	"bocchikitsunei_webportfolio/internal/repository"
	"bocchikitsunei_webportfolio/internal/utils/v"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type userService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewUserService(userRepo repository.UserRepository, jwtSecret string) userService {
	return userService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
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

func (s userService) GetUserId(userid int) (*entities.User, error) {
	user, err := s.userRepo.GetUserById(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := entities.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
	return &userResponse, nil
}

func (s userService) GetUserParams(userid int) (*entities.User, error) {
	user, err := s.userRepo.GetUserByParams(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := entities.User{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}
	return &userResponse, nil
}

////////////////////////////////////////////////////////////////////////////////////

func (s userService) GetEditUserProfile(userid int) (*entities.User, error) {
	user, err := s.userRepo.GetEditUserProfile(userid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userResponse := entities.User{
		UserID:   user.UserID,
		Username: user.Username,
	}
	return &userResponse, nil
}

func (s userService) UpdateEditUserProfile(userid int, req dtos.EditUserProfileRequest) (*entities.User, error) {
	user := &entities.User{
		UserID:   v.UintPtr(userid),
		Username: req.Username,
	}

	err := s.userRepo.UpdateEditUserProfile(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (s userService) Register(request dtos.RegisterRequest) (*dtos.UserInfoResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(v.ByteSlice(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := entities.User{
		Username: request.Username,
		Password: v.Ptr(string(hashedPassword)),
	}

	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dtos.UserInfoResponse{
		UserID:   user.UserID,
		Username: user.Username,
	}, nil

}

func (s userService) Login(request dtos.LoginRequest, jwtSecret string) (*dtos.LoginResponse, error) {
	username := *request.Username

	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid username or password")
		}
		return nil, err
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword(v.ByteSlice(user.Password), []byte(*request.Password)); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	// Generate JWT token
	claims := jwt.StandardClaims{
		Issuer: strconv.Itoa(int(*user.UserID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Token:    &jwtToken,
	}, nil
}
