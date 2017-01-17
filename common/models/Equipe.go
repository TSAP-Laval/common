package models

import "github.com/jinzhu/gorm"

// Equipe est une modélisation d'une équipe de Joueurs
// dirigiée par un ou plusieurs entraineurs
type Equipe struct {
	gorm.Model
	Nom         string
	Ville       string
	Sport       Sport
	SportID     int
	Niveau      Niveau
	NiveauID    int
	Entraineurs []Entraineur `gorm:"many2many:entraineur_equipe;"`
	Joueurs     []Joueur     `gorm:"many2many:joueur_equipe;"`
}
