package models

type Test struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Questions   []Question `gorm:"constraint:OnDelete:CASCADE"`
	CreatedBy   uint       // ID админа, который создал тест
}

type Question struct {
	ID      uint   `gorm:"primaryKey"`
	TestID  uint   `gorm:"index"`
	Text    string `gorm:"not null"`
	OptionA string
	OptionB string
	OptionC string
	OptionD string
	Answer  string // "A", "B", "C" или "D"
}
