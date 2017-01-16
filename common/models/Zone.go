package models

import "github.com/jinzhu/gorm"

// Zone est une modélisation de la zone de jeu
// de l'action
type Zone struct {
	gorm.Model
	Nom string
}
