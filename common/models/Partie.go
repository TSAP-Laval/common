package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Partie est une modélisation des informations
// sur une partie entre deux équipes
type Partie struct {
	gorm.Model
	EquipeMaison    Equipe
	EquipeMaisonID  int
	EquipeAdverse   Equipe
	EquipeAdverseID int
	Saison          Saison
	SaisonID        int
	Lieu            Lieu
	LieuID          int
	Video           Video
	VideoID         int
	Date            string
}

type playerMatch struct {
	ID           int             `json:"match_id"`
	Date         time.Time       `json:"date"`
	OpposingTeam string          `json:"opposing"`
	Metrics      []displayMetric `json:"metrics"`
}
