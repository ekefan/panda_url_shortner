package database

import (
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func TestMain(m *testing.M) {
    dbConn, err := gorm.Open(sqlite.Open("./../test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Couldn't connect to database: ", err)
    }

    ts := NewStore(dbConn)
    if err := ts.RunMigrations(dbConn); err != nil {
        log.Fatal("Failed to run migrations: ", err)
    }

    // Run tests
    exitCode := m.Run()

    // Close database connection
    sqlDB, _ := dbConn.DB()
    sqlDB.Close()

    // Exit with the appropriate exit code
    os.Exit(exitCode)
}

