package models

import "github.com/jinzhu/gorm"

// Niveau est une modélisation du niveau d'un joueur
type Niveau struct {
	gorm.Model
	Nom string
}
