package migrations

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key;autoIncrement"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string
	CreatedAt time.Time  `gorm:"autoCreateTime; not null"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime; not null"`
	DeletedAt *time.Time `gorm:"index"`
}
