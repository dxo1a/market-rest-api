package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Phone    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Email    string
	FullName string
	Address  string
}
