package database

import (
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


//createURLArgs: struct for holding args for creating a new URL-row
type CreateURLArgs struct {
	ShortCode string
	LongURL string
}

type GetURLArgs struct {
	ShortCode string
}


type Store interface{
	CreateURL(args CreateURLArgs) (URL, error)
	GetURL(args GetURLArgs) (URL, error)
	RunMigrations(db *gorm.DB) error

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
func (s *Query) CreateURL(args CreateURLArgs) (URL, error){
	url_row := URL{ShortCode: args.ShortCode, LongURL: args.LongURL}
	result := s.db.Create(&url_row)
	return url_row, result.Error
}

//GetURL: makes a query to the database and returns the URL with short_code specified in args
func (s *Query) GetURL(args GetURLArgs) (URL, error){ 
	urlRow := URL{}
	result := s.db.Where("short_code = ?", args.ShortCode).First(&urlRow)
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
