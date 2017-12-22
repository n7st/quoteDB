package model

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model

	Name string

	Heads []Head
}
