package util

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/n7st/quoteDB/model"
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
		&model.Head{},
		&model.Line{},
	)
}
