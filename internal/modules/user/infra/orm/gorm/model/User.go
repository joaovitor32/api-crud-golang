package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name string // A regular string field
	Age int
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
