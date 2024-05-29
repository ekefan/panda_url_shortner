package database

import (
	// "gorm.io/driver/sqlite"
	"fmt"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

// RunMigrations runs database migrations if the table url don't exist yet
func (s *Store) RunMigrations() {
	// check if table exists else migrate schema/model
	check := !s.db.Migrator().HasTable(&URL{})
	fmt.Println(check)
	if check {
		s.db.AutoMigrate(&URL{})
	}
}

func (s *Store) CreateURL(args createURLArgs) *URL{
	url_row := URL{ShortCode: args.shortCode, LongURL: args.longURL}
	s.db.Create(&url_row)
	fmt.Println("From line 32: Printing new_url_row", url_row)
	return &url_row
}
func (s *Store) GetURL() (url URL)          { return }
func (s *Store) UpdateURL() (url URL)       { return }
func (s *Store) DeleteURl()                 {}
