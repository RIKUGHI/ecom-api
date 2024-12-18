package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/rikughi/ecom-api/internal/entity"
	"github.com/rikughi/ecom-api/internal/model"
	"github.com/rikughi/ecom-api/internal/model/converter"
	"github.com/rikughi/ecom-api/internal/repository"
	"github.com/rikughi/ecom-api/internal/util"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) Create(c *gin.Context, request *model.RegisterUserRequest) (response *model.UserResponse, e error) {
	tx := s.DB.WithContext(c).Begin()
	defer func() {
		if message := recover(); message != nil {
			tx.Rollback()
			s.Log.Warnf("Rollback transaction: %+v", message)
			e = util.ErrInternalServer
		} else if err := tx.Commit().Error; err != nil {
			s.Log.Warnf("Failed to commit transaction: %+v", err)
			e = util.ErrInternalServer
		}
	}()

	userEmail, err := s.UserRepository.FindByEmail(tx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Log.Warnf("Failed to get email user from database : %+v", err)
		return nil, util.ErrInternalServer
	}

	if userEmail != nil {
		s.Log.Warnf("User already exists : %+v", err)
		return nil, util.ErrUserExists
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, util.ErrPasswordHash
	}

	user := &entity.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  string(password),
	}

	if err := s.UserRepository.Create(tx, user); err != nil {
		s.Log.Warnf("Failed create user to database : %+v", err)
		return nil, util.ErrCreateUser
	}

	return converter.UserToResponse(user), nil
}

func (s *UserService) Login(c *gin.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	user, err := s.UserRepository.FindByEmail(s.DB, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Log.Warnf("Failed to get email user from database : %+v", err)
		return nil, util.ErrInternalServer
	}

	if user == nil {
		s.Log.Warnf("User not found : %+v", err)
		return nil, util.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		s.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, util.ErrInvalidCreds
	}

	token, err := util.CreateJWT("secret", user.ID)
	if err != nil {
		s.Log.Warnf("Failed to create jwt token : %+v", err)
		return nil, util.ErrInternalServer
	}

	user.Token = token

	return converter.UserToResponse(user), nil
}

func (s *UserService) Current(c *gin.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	user := new(entity.User)

	err := s.UserRepository.FindById(s.DB, user, request.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Log.Warnf("Failed to get user from database : %+v", err)
		return nil, util.ErrInternalServer
	}

	if user.ID == 0 {
		s.Log.Warnf("User not found : %+v", err)
		return nil, util.ErrUserNotFound
	}

	return converter.UserToResponse(user), nil
}
