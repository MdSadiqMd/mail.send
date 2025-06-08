package services

import (
	"errors"

	"github.com/MdSadiqMd/mail.send/internal/middleware"
	"github.com/MdSadiqMd/mail.send/internal/models"
	"github.com/MdSadiqMd/mail.send/internal/repository"
	"github.com/MdSadiqMd/mail.send/pkg/config"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
)

type UserService struct {
	UserRepo repository.UserRepository
	Auth     middleware.Auth
	Config   config.AppConfig
}

var serviceLogger = logger.New("UserService")

func (s UserService) Signup(input models.User) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		serviceLogger.Error("failed to hash password: %v", err)
		return "", errors.New("failed to hash password")
	}

	user, err := s.UserRepo.CreateUser(&models.User{
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Email:      input.Email,
		Password:   hashedPassword,
		Role:       input.Role,
		IsVerified: false,
	})
	if err != nil {
		serviceLogger.Error("error in creating user: %v", err)
		return "", errors.New("failed to create user")
	}

	return s.Auth.GenerateToken(user.Id, user.Email, user.Role)
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.FindUserByEmail(email)
	if err != nil {
		serviceLogger.Error("error in finding user by email: %v", err)
		return "", errors.New("failed to find user")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		serviceLogger.Error("password does not match: %v", err)
		return "", errors.New("password does not match")
	}

	return s.Auth.GenerateToken(user.Id, user.Email, user.Role)
}

func (s UserService) FindUserByEmail(email string) (*models.User, error) {
	user, err := s.UserRepo.FindUser(email)
	if err != nil {
		serviceLogger.Error("error in finding user: %v", err)
		return nil, errors.New("failed to find user")
	}
	return user, nil
}

func (s UserService) IsVerifiedUser(id uint) bool {
	user, err := s.UserRepo.FindUserById(id)
	if err != nil {
		serviceLogger.Error("error in verifying user: %v", err)
		return false
	}
	return user.IsVerified
}
