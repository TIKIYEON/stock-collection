package models

// Stock table containing the stock id
type Stock struct {
	SID       uint        `gorm:"primaryKey;column:sid"`
	Portfolio []Portfolio `gorm:"many2many:portfolio_stocks;" json:"portfolios"`
}

func (Stock) TableName() string {
	return "stocks"
}
