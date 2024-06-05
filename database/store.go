package database

import "gorm.io/gorm"


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