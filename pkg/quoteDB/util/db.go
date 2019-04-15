// The util package provides functionality needed by all binaries in the
// project.
package util

import (
	"log"

	"git.netsplit.uk/mike/quoteDB/pkg/quoteDB/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB() creates a database connection.
func InitDB(config *Config) *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DBPath)

	if err != nil {
		log.Fatal(err)
	}

	autoMigrate(db)

	return db
}

// autoMigrate() initialises models in the list as database tables
func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Channel{},
		&model.Head{},
		&model.Line{},
	)
}
