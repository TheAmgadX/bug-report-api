package service

import (
	"errors"

	"github.com/TheAmgadX/bug-report-api/internals/models"
	"gorm.io/gorm"
)

type BugService struct {
	db *gorm.DB
}

func (s *BugService) CreateBugReport(bug *models.Bug) error {
	result := s.db.Create(&bug)

	return result.Error
}

func (s *BugService) UpdateBugReport(bug *models.Bug) error {
	return s.db.Model(&models.Bug{}).Select("title", "description", "severity", "status").
		Where("id = ?", bug.ID).Updates(&bug).Error
}

func (s *BugService) UpdateBugsReports(bugs []*models.Bug) error {
	// one transaction if one fails all fail.
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, b := range bugs {
			err := tx.Model(&models.Bug{}).Select("title", "description", "severity", "status").
				Where("id = ?", b.ID).Updates(&b).Error

			if err != nil {
				return err
			}

		}

		return nil
	})
}

func (s *BugService) DeleteBugReport(bugID int) error {
	if bugID == 0 {
		return errors.New("cannot delete bug: ID is 0")
	}
	
	return s.db.Delete(&models.Bug{ID: bugID}).Error
}

func (s *BugService) DeleteBugsReports(bugs []*models.Bug) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, b := range bugs {
			if b.ID == 0 {
				return errors.New("cannot delete bug with ID: 0")
			}
			if err := tx.Delete(&b).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *BugService) GetBugsReportsByUser(userID int) (bugs []*models.Bug, err error) {
	err = s.db.Model(&models.Bug{}).Where("user_id = ?", userID).Find(&bugs).Error
	if err != nil {
		return nil, err
	}

	return bugs, nil
}

func (s *BugService) GetAllBugsReports() (bugs []*models.Bug, err error) {
	err = s.db.Model(&models.Bug{}).Find(&bugs).Error

	if err != nil {
		return nil, err
	}

	return bugs, nil
}

func (s *BugService) GetBugsReports(filter *BugFilter) ([]*models.Bug, error) {

	// method chaining so we can mix filterations and easly add new filters
	query := s.db.Model(&models.Bug{})

	if filter.UserID != nil {
		query = query.Where("user_id = ?", *filter.UserID)
	}

	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	if filter.Severity != nil {
		query = query.Where("severity = ?", *filter.Severity)
	}

	if filter.CreatedAt != nil {
		query = query.Where("created_at = ?", *filter.CreatedAt)
	}

	bugs := []*models.Bug{}

	err := query.Find(&bugs).Error

	if err != nil {
		return nil, err
	}

	return bugs, nil
}
