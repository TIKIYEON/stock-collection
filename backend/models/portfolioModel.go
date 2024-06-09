package models

type Portfolio struct {
	PortfolioID uint    `gorm:"primaryKey;column:pid" json:"pid"`
	UserID      uint    `gorm:"column:user_id" json:"user_id"`
	Stocks      []Stock `gorm:"many2many:portfolio_stocks" json:"stocks"`
}
