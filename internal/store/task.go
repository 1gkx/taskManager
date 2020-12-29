package store

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Description string
	Status      string
	Type        string
	Author      User
	Executor    User
	Priority    int
	Start       time.Time
	Deadline    time.Time
}

// Create
