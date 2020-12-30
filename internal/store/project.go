package store

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Manager     string `json:"manager"`
	Tasks       string `json:"tasks"`
	Members     string `json:"members"`
}
