package models

import (
)

// Stock table containing the stock id
type Stock struct {
    SID uint `gorm:"primaryKey;column:sid"`
}

