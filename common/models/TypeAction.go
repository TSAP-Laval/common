package models

import "github.com/jinzhu/gorm"

// TypeAction est une modélisation des types
// possibles  d'action
type TypeAction struct {
	gorm.Model
	Nom string `gorm:"unique"`
}
