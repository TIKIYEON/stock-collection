package models

// Stock table containing the stock id
type Portfolio struct {
	PID    uint    `gorm:"primaryKey;column:pid" json:"pid"`
	Value  float64 `gorm:"type:decimal(10,2)" json:"value"`
	UserID uint    `gorm:"column:user_id" json:"user_id"`
	User   User    `gorm:"foreignKey:UserID;references:UID"`
}
