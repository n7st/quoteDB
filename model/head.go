// Package model contains the quote database's schema files.
package model

import "github.com/jinzhu/gorm"

// Head{} contains generic information for quotes and has a one-to-many
// relationship with Line{}s against it.
type Head struct {
	gorm.Model

	Channel string
	Title   string

	Lines []Line
}
