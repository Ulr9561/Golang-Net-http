package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	Desc string `json:"desc"`
	Done bool   `json:"done"`
}
