package service

import (
	"errors"

	"github.com/TheAmgadX/bug-report-api/internals/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func (s *UserService) CreateUser(user *models.User) error {
	// user validation will be in the handlers not here.
	result := s.db.Create(&user)

	return result.Error
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user *models.User

	if err := s.db.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	if err := s.db.Model(&models.User{}).Where("email = ?", email).First(&user).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(userID int) error {
	return s.db.Delete(&models.User{ID: userID}).Error
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.db.Model(&models.User{}).Where("id = ?", user.ID).
		Select("username", "password", "email").Updates(&user).Error
}
