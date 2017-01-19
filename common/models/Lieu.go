package models

import "github.com/jinzhu/gorm"

// Lieu est une modélisation des endroits possibles
// pour un match
type Lieu struct {
	gorm.Model
	Nom     string
	Ville   string
	Adresse string
}
