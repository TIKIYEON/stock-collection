package models

type Portfoliostocks struct {
	PortfolioID uint      `gorm:"primaryKey; column:portfolioid" json:"portfolio_id"`
	StockID     uint      `gorm:"primaryKey; column:stockid" json:"stock_id"`
	Portfolio   Portfolio `gorm:"foreignKey:PortfolioID;references:PID;constraint:OnDelete:CASCADE;"`
	Stock       Stock     `gorm:"foreignKey:StockID;references:SID;constraint:OnDelete:CASCADE;"`
}

func (Portfoliostocks) TableName() string {
	return "portfoliostocks"
}
