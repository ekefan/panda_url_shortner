package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func (s *Query) RunMigrations(db *gorm.DB) error {
	hasUser := s.db.Migrator().HasTable(&URL{})
	hasURL := s.db.Migrator().HasTable(&USER{})
	if hasUser && hasURL {
		return nil
	}
    sqlDB, err := db.DB()
    if err != nil {
        return fmt.Errorf("could not get database instance: %v", err)
    }

    driver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
    if err != nil {
     return fmt.Errorf("could not create migration driver: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://database/migrations",
        "sqlite", driver)
    if err != nil {
        return fmt.Errorf("could not create migrate instance: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("could not apply migrations: %v", err)
    }
	// if err := m.Down(); err != nil && err == migrate.ErrNoChange{
	// 	return fmt.Errorf("could not rollback migrations: %v", err)
	// }
	return nil
}