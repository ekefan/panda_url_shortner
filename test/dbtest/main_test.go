package dbtest

import (
	"log"
	"os"
	"testing"

	db "github.com/ekefan/panda_url_shortner/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ts db.Store

func TestMain(m *testing.M) {
    dbConn, err := gorm.Open(sqlite.Open("./../test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Couldn't connect to database: ", err)
    }

    ts = db.NewStore(dbConn)
    if err := ts.RunMigrations(dbConn, 1); err != nil {
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