package models

import "github.com/jinzhu/gorm"

// Position est une modélisation du nom de la Position
// du joueur
type Position struct {
	gorm.Model
	Nom string `gorm:"unique"`
}
