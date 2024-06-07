package models

import (
	"time"
)

// StockElement table containing all the stock elements from the data sourc
type Portfolioelement struct {
	BuyingDate  time.Time `gorm:"type:date"`
	BuyingPrice float64   `gorm:"type:decimal(10,4)"`
	PriceChange float64   `gorm:"type:decimal(10,4)"`
	PortfolioID uint      `gorm:"column:portfolio_id"`
	StockID     uint      `gorm:"column:stock_id"`
	Portfolio   Portfolio `gorm:"foreignKey:PortfolioID;references:PID;"`
	Stock       Stock     `gorm:"foreignKey:StockID;references:SID;"`
}
