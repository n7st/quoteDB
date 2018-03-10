// The model package contains the quote database's schema files.
package model

import "github.com/jinzhu/gorm"

// Channel{} describes an IRC channel which has many quotes ([]Head{}) stored
// against it.
type Channel struct {
	gorm.Model

	Name    string
	Deleted bool

	Heads []Head
}
