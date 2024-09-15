package entity

import (
	"time"
)

// User structure to describe user table
type User struct {
	UserId      string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string     `gorm:"type:varchar(100);not null"`
	Age         int        `gorm:"not null; index:idx_user_age"`
	Gender      string     `gorm:"type:varchar(10); index:idx_user_gender; not null"`
	Location    string     `gorm:"type:geography(POINT,4326); index:idx_user_location"`
	Interests   []string   `gorm:"type:text[]"`
	Preferences Preference `gorm:"foreignKey:UserID;references:UserId"`
	LastActive  time.Time  `gorm:"not null; index:idx_user_last_active"`
	Rank        int        `gorm:"-"`
}

func (User) TableName() string {
	return "public.user"
}
