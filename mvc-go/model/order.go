package model

import (
	"time"
)

type Order struct {
	Id          int       `gorm:"primaryKey"`
	Monto_Final float32   `gorm:"type:decimal(60,4);not null"`
	Fecha       time.Time `gorm:"not null"`
	Id_User     int       `gorm:"type:integer;not null"`
}

type Orders []Order
