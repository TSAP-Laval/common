package models

import "github.com/jinzhu/gorm"

// PlayerStats représente les statistiques d'un joueur
type PlayerStats struct {
	ID        uint          `json:"player_id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Team      displayTeam   `json:"team"`
	Season    displaySeason `json:"season"`
	Matches   []playerMatch `json:"matches"`
}

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
