package models

type User struct {
	UID         uint       `gorm:"primaryKey;column:uid" json:"uid"`
	Password    string     `gorm:"not null" json:"password"`
	Mail        string     `gorm:"not null" json:"mail"`
	PhoneNumber string     `json:"phone_number"`
	Portfolio   *Portfolio `gorm:"foreignKey:UserID;references:UID" json:"portfolio"`
}

func (User) TableName() string {
	return "users"
}
