package database

import (
	"time"
	// "gorm.io/gorm"
)


// URL: Model for the handling short codes and full urls
type URL struct {
	ID        uint      `gorm:"primaryKey; autoIncrement; not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	ShortCode string    `json:"short_code" gorm:"not null; unique"`
	LongURL   string    `json:"long_url" gorm:"not null"`
}


