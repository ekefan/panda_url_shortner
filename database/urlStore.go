package database

import (
	// "errors"

	"fmt"
	"time"

	"gorm.io/gorm"
)

// const minShortCodeLen int = 4

// createURLArgs: struct for holding args for creating a new URL-row
type CreateURLArgs struct {
	Owner     string `json:"owner"`
	ShortCode string `json:"short_code"`
	LongURL   string `json:"long_url"`
}

type GetURLArgs struct {
	ShortCode string `json:"short_code"`
}

// RunMigrations runs database migrations if the table url don't exist yet
func (s *Query) CreateURL(args CreateURLArgs) (URL, error) {
	url_row := URL{Owner: args.Owner, ShortCode: args.ShortCode, LongURL: args.LongURL}
	result := s.db.Create(&url_row)
	return url_row, result.Error
}

// GetURL: makes a query to the database and returns the URL with short_code specified in args
func (s *Query) GetURL(args GetURLArgs) (URL, error) {
	urlRow := URL{}
	result := s.db.Where("short_code = ?", args.ShortCode).First(&urlRow)
	return urlRow, result.Error
}

// Code from yesterday

// GetURL
func (s *Query) getUrlForUpdate(shortCode string) (URL, error) {
	urlRow := URL{}
	result := s.db.Where("owner = ?", shortCode).First(&urlRow)
	return urlRow, result.Error

}

type UpdateShortCodeArgs struct {
	Owner     string `json:"owner"`
	ShortCode string `json:"short_code"`
}

// UpdateShortCode a transaction to update the short code off the url
func (s *Query) UpdateShortCode(args UpdateShortCodeArgs) (URL, error) {
	urlRow := URL{}
	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&urlRow).Where("owner = ?", args.Owner).
			Updates(URL{ShortCode: args.ShortCode, UpdatedAt: time.Now()})

		// check if the update query was successful
		if result.Error != nil {
			return result.Error
		}

		//check if the update query affected at least on row
		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows affected, possibly invalid owner: %v", args.Owner)
		}

		return nil
	})

	if txErr != nil {
		return urlRow, txErr
	}

	//get the newly update url row by the newShort code
	urlRow, err := s.getUrlForUpdate(args.ShortCode)
	if err != nil {
		return urlRow, err
	}

	return urlRow, nil
}

// func (s *Store) DeleteURl() \begin, get the user delete the url based on the url id{}
//=====edit URL model to contain title and user generated shortcode
//UPDATEURL should edit the title of the url, and the shortcode.
