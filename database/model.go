package database

import (
	"time"
	// "gorm.io/gorm"
)


// URL: Model for the handling short codes and full urls
type URL struct {
	Owner        uint
	ShortCode string 
	LongURL   string 
	CreatedAt time.Time
	UpdatedAt time.Time

}

type USER struct {
	ID uint
	Name string
	Email string
	Password string
	CreatedAt string
}
