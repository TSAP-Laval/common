package models

import "github.com/jinzhu/gorm"

// Sport est une modélisation des noms possibles
// des sports pratiqués
type Sport struct {
	gorm.Model
	Nom string
}
