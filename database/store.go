package database

import (
	// "gorm.io/driver/sqlite"
	"fmt"
	"gorm.io/gorm"
)

type Store interface{
	CreateURL(args createURLArgs) (URL, error)
	GetURL(args getURLArgs) (URL, error)
	RunMigrations() error

}

type Query struct {
	db *gorm.DB
}


func NewStore(db *gorm.DB) Store {
	return &Query{
		db: db,
	}
}

// RunMigrations runs database migrations if the table url don't exist yet
func (s *Query) RunMigrations() error {
	// check if table exists else migrate schema/model
	check := !s.db.Migrator().HasTable(&URL{})
	fmt.Println(check)
	if check {
		err := s.db.AutoMigrate(&URL{})
		if err != nil {
			return fmt.Errorf("could not migrate database: %s", err)
		}
	}
	return nil
}

func (s *Query) CreateURL(args createURLArgs) (URL, error){
	url_row := URL{ShortCode: args.shortCode, LongURL: args.longURL}
	result := s.db.Create(&url_row)
	return url_row, result.Error
}

//GetURL: makes a query to the database and returns the URL with short_code specified in args
func (s *Query) GetURL(args getURLArgs) (URL, error){ 
	urlRow := URL{}
	result := s.db.Where("short_code = ?", args.shortCode).First(&urlRow)
	return urlRow, result.Error
}

//posible upgrade.. getURL only where the deletedAt is zerovalue

// TODO: implement updateURL and DeleteURL when You implement USER model
// func (s *Store) UpdateURL() (url URL)       { return }
// func (s *Store) DeleteURl()                 {}




// TODO: Implement userStore.go
//CreateUser
//Login
//=====edit URL model to contain title and user generated shortcode
//UPDATEURL should edit the title of the url, and the shortcode.
