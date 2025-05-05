package models

import (
	"time"

	"gorm.io/gorm"
)

type BugSeverity int8
type BugStatus int8

const (
	High BugSeverity = iota + 1
	Medium
	Low
)

const (
	New BugStatus = iota + 1
	InProgress
	Solved
)

func (s *BugStatus) ToString() string {
	switch *s {
	case New:
		return "New"
	case InProgress:
		return "In Progress"
	case Solved:
		return "Solved"
	default:
		return "UnKnown Value"
	}
}

func (s *BugSeverity) ToString() string {
	switch *s {
	case High:
		return "High"
	case Medium:
		return "Medium"
	case Low:
		return "Low"
	default:
		return "UnKnown Value"
	}
}

type Bug struct {
	ID          int         `gorm:"primaryKey;autoIncrement"`
	Title       string      `gorm:"not null" json:"title"`
	Description string      `json:"description"`
	Severity    BugSeverity `gorm:"index" json:"severity"`
	Status      BugStatus   `gorm:"index"`

	UserID int `gorm:"not null" json:"user_id"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt
}
