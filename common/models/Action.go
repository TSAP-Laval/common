package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Action est une modélisation des informations sur une
// action exécutée par un joueur
type Action struct {
	gorm.Model
	TypeAction      TypeAction
	ActionPositive  bool
	Zone            Zone
	ZoneID          int
	Partie          Partie
	PartieID        int
	X1              float64
	Y1              float64
	X2              float64
	Y2              float64
	Temps           time.Duration
	PointageMaison  int
	PointageAdverse int
	Joueur          Joueur
	JoueurID        int
}
