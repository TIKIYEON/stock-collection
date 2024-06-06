package models

type User struct {
    UID uint `gorm:"primaryKey" json:"uid"`
    Password string `gorm:"not null" json:"password"`
    Mail string `gorm:"not null" json:"mail"`
    PhoneNumber string `json:"phone_number"`
}
