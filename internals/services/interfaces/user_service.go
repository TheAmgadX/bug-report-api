package interfaces

import (
	"github.com/TheAmgadX/bug-report-api/internals/models"
	service "github.com/TheAmgadX/bug-report-api/internals/services"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	DeleteUser(userID int) error
	UpdateUser(user *models.User) error
}

var _ UserService = (*service.UserService)(nil) // to check if the struct implements the interface or not.
