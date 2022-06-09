package model

type User struct {
	Id       int    `gorm:"primaryKey"`                 
	Name     string `gorm:"type:varchar(350);not null"` 
	Last_Name string `gorm:"type:varchar(250);not null"`
	User_Name string `gorm:"type:varchar(150);not null"` 
	Password string `gorm:"type:varchar(150);not null"` 
}

type Users []User
