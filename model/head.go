package model

import "github.com/jinzhu/gorm"

type Head struct {
	gorm.Model

	Channel string
	Title   string
	Lines   []Line
}