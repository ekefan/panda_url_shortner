package database

import (
	// "gorm.io/driver/sqlite"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// URL: Model for the handling short codes and full urls
type URL struct {
	ID        uint      `gorm:"autoIncrement; not null"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
	ShortCode string    `json:"short_code" gorm:"not null"`
	LongURL   string    `json:"long_url" gorm:"not null"`
}

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

func (s *Store) CreateURL(shortCode string) {}
func (s *Store) GetURL() (url URL)          { return }
func (s *Store) UpdateURL() (url URL)       { return }
func (s *Store) DeleteURl()                 {}
