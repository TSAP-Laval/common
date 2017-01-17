package models

import "github.com/jinzhu/gorm"

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
	Actions         []Action
}

// Expand effectue un fetch de tous les children de la partie
// (has-many, has-one, pas belongs-to)
func (p *Partie) Expand(db *gorm.DB) {
	db.Model(p).Related(&(p.Actions))
	db.Model(p).Related(&(p.EquipeMaison), "EquipeMaisonID")
	db.Model(p).Related(&(p.EquipeAdverse), "EquipeAdverseID")
}
