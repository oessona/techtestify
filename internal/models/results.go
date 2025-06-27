package models

import "time"

type Result struct {
	ID      uint `gorm:"primaryKey"`
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
	TestID  uint
	Test    Test `gorm:"foreignKey:TestID"`
	Score   int
	Total   int
	Created time.Time
}
