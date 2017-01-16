package models

import "github.com/jinzhu/gorm"

// Metrique est la modélisation d'une unité de calcul
type Metrique struct {
	gorm.Model
	Nom      string `gorm:"primary_key"`
	Equation string `gorm:"primary_key"`
	Equipe   Equipe
	EquipeID int
}

type displayMetric struct {
	DisplayPair
	Value     float64 `json:"value"`
	Deviation float64 `json:"value"`
}
