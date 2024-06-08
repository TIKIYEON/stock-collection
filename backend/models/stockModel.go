package models

type Stock struct {
	SID        uint        `gorm:"primaryKey;column:sid"`
	Portfolios []Portfolio `gorm:"many2many:portfolio_stocks;"`
}

func (Stock) TableName() string {
	return "stocks"
}
