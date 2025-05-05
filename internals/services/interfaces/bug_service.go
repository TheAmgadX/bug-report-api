package interfaces

import (
	"github.com/TheAmgadX/bug-report-api/internals/models"
	"github.com/TheAmgadX/bug-report-api/internals/services"
)

type BugService interface {
	CreateBugReport(*models.Bug) error
	UpdateBugReport(*models.Bug) error
	UpdateBugsReports([]*models.Bug) error
	DeleteBugReport(bugID int) error
	DeleteBugsReports([]*models.Bug) error
	GetBugsReportsByUser(int) ([]*models.Bug, error)
	GetAllBugsReports() ([]*models.Bug, error)
	GetBugsReports(filter *service.BugFilter) ([]*models.Bug, error)
}

var _ BugService = (*service.BugService)(nil) // to check if the struct implements the interface.

