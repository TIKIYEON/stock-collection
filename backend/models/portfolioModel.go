package models

// Stock table containing the stock id
type Portfolio struct {
	PID    uint    `gorm:"primaryKey;column:sid"`
	Value  float64 `gorm:"type:decimal(20,16)"`
	UserID uint    `gorm:"column:uid"`
	User   User    `gorm:"foreignKey:UserID;references:UID;constrant:OnDelete:CASCADE;"`
}
