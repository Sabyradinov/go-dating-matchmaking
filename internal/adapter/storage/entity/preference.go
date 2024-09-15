package entity

// Preference structure to describe user's preference
type Preference struct {
	ID          int     `gorm:"primaryKey"`
	UserID      string  `gorm:"type:uuid;not null"`
	Gender      string  `gorm:"type:varchar(10)"`
	AgeRangeMin int     `gorm:"type:int"`
	AgeRangeMax int     `gorm:"type:int"`
	MaxDistance float64 `gorm:"type:float"`
}

func (Preference) TableName() string {
	return "public.preference"
}
