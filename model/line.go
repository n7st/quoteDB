package model

import "github.com/jinzhu/gorm"

type Line struct {
	gorm.Model

	Content string
	Author  string

	Head   Head
	HeadID uint
}
