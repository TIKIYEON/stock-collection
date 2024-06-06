package models

type User struct {
    UID uint `gorm:"primaryKey"`
    Password string `gorm:"not null"`
    Mail string `gorm:"not null"`
    PhoneNumber string
}
