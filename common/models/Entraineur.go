package models

import (
	"github.com/jinzhu/gorm"
)

// Entraineur est une modelisation d'un entraineur
// comprenant son id, son nom, son prenom, son email
// et ses équipes
type Entraineur struct {
	gorm.Model
	Prenom   string
	Nom      string
	Email    string
	PassHash string
	Token    string
	Equipes  []Equipe `gorm:"many2many:entraineur_equipe;"`
}
