package database

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func TestMain(m *testing.M){
	db, err := gorm.Open(sqlite.Open("./../test.db"), &gorm.Config{
	})
	if err != nil {
		log.Fatal("Couldn't connect to database: ", err)
	}

	ts = NewStore(db)
	ts.RunMigrations()
	os.Exit(m.Run())
	
}