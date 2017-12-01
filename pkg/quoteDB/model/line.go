// Package model contains the quote database's schema files.
package model

import "github.com/jinzhu/gorm"

// Line{} has a many-to-one relationship with Head{}. It contains information
// about a one-line part of a quote.
type Line struct {
	gorm.Model

	Content string
	Author  string

	Head   Head
	HeadID uint
}
