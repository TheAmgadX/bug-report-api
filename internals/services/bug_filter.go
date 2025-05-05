package service

import (
	"time"

	"github.com/TheAmgadX/bug-report-api/internals/models"
)

type BugFilter struct {
	UserID    *int
	Severity  *models.BugSeverity
	Status    *models.BugStatus
	CreatedAt *time.Time
}
