package models

import "github.com/jinzhu/gorm"

// Video est la modélisation des informations
// sur une Video
type Video struct {
	gorm.Model
	Path           string
	AnalyseTermine bool
}
