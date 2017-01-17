package models

import "github.com/jinzhu/gorm"

// Saison est un modèlisation d'une saison
// sportive, comprenant son ID et ses années
type Saison struct {
	gorm.Model
	Annees string `gorm:"size:10"`
}
