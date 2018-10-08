package shared

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres provider
	"github.com/pauliushchyk/test-go-webapp/config"
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
