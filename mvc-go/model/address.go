package model

type Address struct {
	Id           int    `gorm:"primaryKey"`
	Neighborhood string `gorm:"type:varchar(250);not null"`
	Street       string `gorm:"type:varchar(150);not null"`
	Number       int    `gorm:"type:int(150);not null"`
	Id_User      int    `gorm:"type:integer;not null"`
}

type Addresses []Address
