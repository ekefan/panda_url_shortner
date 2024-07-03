package database

import "gorm.io/gorm"

type Store interface {
	CreateURL(args CreateURLArgs) (URL, error)
	GetURL(args GetURLArgs) (URL, error)
	RunMigrations(db *gorm.DB, flag int) error
	CreateUser(args CreateUserArgs) (USER, error)
	GetUser(args GetUserArgs) (USER, error)
	GetURLs(args GetURLsArg) ([]URL, error)
	TxUpdateShortCode(args TxUrlArgs) (URL, error)
	TxUpdateUser(args TxUserArgs) (USER, error)
	TxDeleteUser(args TxUserArgs) error
	TxDeleteUrl(args TxUrlArgs) error 
}

type Query struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) Store {
	return &Query{
		db: db,
	}
}
