package models

import "github.com/jinzhu/gorm"

// Administrateur est une modélisation des informations
// de connexion d'un Administrateur
type Administrateur struct {
	gorm.Model
	Email                 string
	PassHash              string
	TokenReinitialisation string
	TokenConnexion        string
}
