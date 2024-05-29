package main

import (
	"github.com/ekefan/panda_url_shortner/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "time"
)


// main: the main entry point for the url_shorten server
func main() {
	//create a connection the the sqlite database file using gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
	})
	if err != nil {
		panic("failed to connect database")
	}
	//Create an instance of a store that connects to the database
	s := database.NewStore(db)

	// run migrations create database tables if they don't exist in the db
	s.RunMigrations()
	/*
		Couldn't separated migration process from the main code using gorm
		db.Migrator().DropTable(&database.URL{}) //so I used this till I got the
												// schema I wanted
	*/
	// Start Server


}
