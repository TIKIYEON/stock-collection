package models

// Stock table containing the stock id
type Stock struct {
	SID uint `gorm:"primaryKey;column:sid"`
}

func (Stock) TableName() string {
	return "stocks"
}
