package database

import (
	"time"
	// "gorm.io/gorm"
)

// URL: Model for the handling short codes and full urls
type URL struct {
	Owner     uint      `json:"owner" gorm:"primaryKey;not null"`
	ShortCode string    `json:"short_code" gorm:"not null;unique"`
	LongURL   string    `json:"long_url" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:0"`
}

type USER struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name      string `json:"name" gorm:"not null;index"`
	Email     string `json:"email" gorm:"not null;unique"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
}
