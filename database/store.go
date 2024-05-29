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

func (s *Store) CreateURL(args createURLArgs) (URL, error){
	url_row := URL{ShortCode: args.shortCode, LongURL: args.longURL}
	result := s.db.Create(&url_row)
	return url_row, result.Error
}

//GetURL: makes a query to the database and returns the URL with short_code specified in args
func (s *Store) GetURL(args getURLArgs) (URL, error){ 
	urlRow := URL{}
	result := s.db.Where("short_code = ?", args.shortCode).First(&urlRow)
	return urlRow, result.Error
}//posible upgrade.. getURL only where the deletedAt is zerovalue
func (s *Store) UpdateURL() (url URL)       { return }
func (s *Store) DeleteURl()                 {}
