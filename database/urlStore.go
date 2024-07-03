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

// GetURLsArg hold fields for getting a list of urls from the database
type GetURLsArg struct {
	Owner  string `json:"owner"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

// GetURLs make a query to database and returns urls from the offset to the limit.
func (s *Query) GetURLs(args GetURLsArg) ([]URL, error) {
	urls := []URL{}
	result := s.db.Limit(args.Limit).Offset(args.Offset).Where("owner = ?", args.Owner).Find(&urls)

	if result.Error != nil {
		return nil, result.Error
	}

	return urls, nil
}

// TxUrlArgs hold fields needed to make an update and delete tx for a url row
type TxUrlArgs struct {
	Owner     string
	CurrentShortCode string
	ShortCode string 
}

// TxUpdateShortCode a transaction to update the short code off the url
// a possible error is shortcode constraint
func (s *Query) TxUpdateShortCode(args TxUrlArgs) (URL, error) {
	urlRow := URL{}
	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&urlRow).Where("owner = ? AND short_code = ?", args.Owner, args.CurrentShortCode).
			Updates(URL{ShortCode: args.ShortCode, UpdatedAt: time.Now()})

		// check if the update query was successful
		if result.Error != nil {
			return result.Error
		}

		//check if the update query affected at least on row
		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}

		return nil
	})

	if txErr != nil {
		return urlRow, txErr
	}

	//get the newly update url row by the newShort code
	arg := GetURLArgs{args.ShortCode}
	urlRow, err := s.GetURL(arg)
	if err != nil {
		return urlRow, err
	}

	return urlRow, nil
}

// TxDeleteUrl a transaction to update the database removing a url row
func (s *Query) TxDeleteUrl(args TxUrlArgs) error {
	urlRow := URL{}
	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		// Find the url row based on the owner
		result := tx.First(&urlRow, "owner = ?", args.Owner)
		if result.Error != nil {
			return fmt.Errorf("user can't delete url: %v", result.Error)
		}

		// Delete the url row based on the shortcode, owner and shortcodes are unique
		result = tx.Where("owner = ? AND short_code = ?", args.Owner, args.ShortCode).Delete(&urlRow)
		if result.Error != nil {
			return result.Error
		}

		// Check if any rows were affected
		if result.RowsAffected == 0 {
			return fmt.Errorf("no rows affected")
		}
		return nil
	})

	if txErr != nil {
		return txErr
	}
	return nil
}

//=====edit URL model to contain title and user generated shortcode
//UPDATEURL should edit the title of the url, and the shortcode.


func (s *Query) GetUrlByOwnerAndShortCode(owner string, shortCode string) (URL, error) {
	var urlRow URL
	result := s.db.Where("owner = ? AND short_code = ?", owner, shortCode).First(&urlRow)
	if result.Error != nil {
		return urlRow, result.Error
	}
	return urlRow, nil
}