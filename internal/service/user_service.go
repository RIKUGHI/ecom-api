package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/model/converter"
	"github.com/rikughi/ecom-api/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	UserRepository *repository.UserRepository
}

func NewUserService(DB *gorm.DB, Log *logrus.Logger, UserRepository *repository.UserRepository) *UserService {
	return &UserService{
		DB:             DB,
		Log:            Log,
		UserRepository: UserRepository,
	}
}

func (s *UserService) Create(c *gin.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	user := &entity.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	}

	if err := s.UserRepository.Create(s.DB, user); err != nil {
		s.Log.Warnf("Failed create user to database : %+v", err)
		return nil, errors.New("Internal Server Error")
	}
	s.Log.Println(user)
	return converter.UserToResponse(user), nil
}
