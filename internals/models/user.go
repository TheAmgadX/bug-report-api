package models

import (
	"time"

	"github.com/TheAmgadX/bug-report-api/internals/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       int  `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"not null;unique;index" json:"username"`
	Email    string `gorm:"not null;unique;index" json:"email"`
	Password string `gorm:"not null" json:"password"`

	Bugs []Bug `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func (u *User) IsValid() bool {
	
	if len(u.Username) < 3 {
		return false
	}
	
	if utils.Valid(u.Email) {
		return false
	}

	if len(u.Password) < 6 {
		return false
	}
	return true
}
