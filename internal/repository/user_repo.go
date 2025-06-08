package repository

import (
	"errors"

	"github.com/MdSadiqMd/mail.send/internal/models"
	logger "github.com/MdSadiqMd/mail.send/pkg/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindUser(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

var repoLogger = logger.New("UserRepository")

func (r userRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		repoLogger.Error("error in creating user: %v", err)
		return &models.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (r *userRepository) FindUser(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).
		First(&user).Error
	if err != nil {
		repoLogger.Error("error in finding user: %v", err)
		return &models.User{}, errors.New("failed to find user")
	}
	return &user, nil
}

func (r *userRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).
	First(&user).Error
	if err != nil {
		repoLogger.Error("error in finding user: %v", err)
		return &models.User{}, errors.New("failed to find user")
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(id uint, user *models.User) (*models.User, error) {
	var updatedUser models.User
	result := r.db.Model(&models.User{}).
		Where("id = ?", id).
		Clauses(clause.Returning{}).
		Updates(user).
		Scan(&updatedUser)
	if result.Error != nil {
		repoLogger.Error("error in updating user: %v", result.Error)
		return &models.User{}, errors.New("failed to update user")
	}
	return &updatedUser, nil
}

func (r *userRepository) DeleteUser(user *models.User) (*models.User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		repoLogger.Error("error in deleting user: %v", err)
		return &models.User{}, errors.New("failed to delete user")
	}
	return user, nil
}
