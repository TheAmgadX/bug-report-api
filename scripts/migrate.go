package scripts

import (
	"github.com/TheAmgadX/bug-report-api/internals/models"
	"gorm.io/gorm"
)

func Migrate(tx *gorm.DB) error {
	tx = tx.Begin()

	// if panic rollback 
	defer func(){
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if !tx.Migrator().HasTable(&models.User{}) {
		if err := tx.Migrator().CreateTable(&models.User{}); err != nil {
			tx.Rollback()
			return err
		}
	}

	if !tx.Migrator().HasTable(&models.Bug{}) {
		if err := tx.Migrator().CreateTable(&models.Bug{}); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}


