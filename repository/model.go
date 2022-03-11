package repository

import (
	"gorm.io/gorm"
)

type Hackaton struct {
	gorm.Model
	Name string
	Devs []Developer `gorm:"foreignkey:HackatonID"`
}

func (Hackaton) TableName() string { return "hackaton" }

type Developer struct {
	gorm.Model
	HackatonID uint
	Position   int
	Name       string
	LastName   string
}

func (Developer) TableName() string { return "developer" }
