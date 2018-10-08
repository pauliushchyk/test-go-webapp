package shared

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres provider
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // sqlite provider
	"github.com/pauliushchyk/test-go-webapp/config"
	"github.com/pauliushchyk/test-go-webapp/models"
)

var db *gorm.DB

// CloseConnection closes connection to database
func CloseConnection() {
	db.Close()
}

// GetConnection returns connection to database
func GetConnection() *gorm.DB {
	return db
}

// OpenConnection opens connection to database
func OpenConnection() *gorm.DB {
	var cs = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", config.Get("db.host"), config.Get("db.port"), config.Get("db.user"), config.Get("db.dbname"), config.Get("db.password"))

	var e error

	db, e = gorm.Open("postgres", cs)

	if e != nil {
		panic("failed to connect database")
	}

	return db
}

// OpenTestConnection opens connection to test database
func OpenTestConnection() *gorm.DB {
	var e error

	db, e = gorm.Open("sqlite3", "/tmp/test.db")

	if e != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Customer{})

	return db
}
