package entity

import (
	"github.com/lib/pq"
	"time"
)

// User structure to describe user table
type User struct {
	UserID     string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string         `gorm:"type:varchar(100);not null"`
	Age        int            `gorm:"not null; index:idx_user_age"`
	Gender     string         `gorm:"type:varchar(10); index:idx_user_gender; not null"`
	Location   string         `gorm:"type:geography(POINT,4326); index:idx_user_location"`
	Interests  pq.StringArray `gorm:"type:text[]"`
	LastActive time.Time      `gorm:"not null; index:idx_user_last_active"`
	Preference *Preference    `gorm:"-"`
	Rank       int            `gorm:"-"`
}

func (User) TableName() string {
	return "public.user"
}
