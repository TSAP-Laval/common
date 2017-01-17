package models

import "github.com/jinzhu/gorm"

// Joueur est une modèlisation d'un joueur
// d'une équipe sportive
type Joueur struct {
	gorm.Model
	Prenom                string
	Nom                   string
	Numero                int
	Email                 string
	PassHash              string
	TokenInvitation       string
	TokenReinitialisation string
	TokenConnexion        string
	Equipes               []Equipe `gorm:"many2many:joueur_equipe;"`
}
