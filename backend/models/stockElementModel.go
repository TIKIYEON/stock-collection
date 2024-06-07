package models

import (
	"time"
)

// StockElement table containing all the stock elements from the data sourc
type StockElement struct {
	Date     time.Time `gorm:"type:date"`
	Open     float64   `gorm:"type:decimal"`
	High     float64   `gorm:"type:decimal(20,16)"`
	Low      float64   `gorm:"type:decimal(20,16)"`
	Close    float64   `gorm:"type:decimal(20,16)"`
	AdjClose float64   `gorm:"type:decimal(20,16)"`
	Volume   int64
	StockID  uint  `gorm:"column:stock_id"`
	Stock    Stock `gorm:"foreignKey:StockID;references:SID;constrant:OnDelete:CASCADE;"`
}
