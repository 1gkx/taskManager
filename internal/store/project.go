package store

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name    string
	Manager string
	Tasks   string
	Members string
}
