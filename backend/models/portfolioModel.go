package models

type Portfolio struct {
	PortfolioID uint    `gorm:"primaryKey;column:pid" json:"pid"`
	Value       float64 `gorm:"type:decimal(10,2)" json:"value"`
	UserID      uint    `gorm:"column:user_id" json:"user_id"`
	Stocks      []Stock `gorm:"many2many:portfolio_stocks" json:"stocks"`
}
