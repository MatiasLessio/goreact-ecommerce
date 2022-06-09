package model

type OrderDetail struct {
	Id              int     `gorm:"primaryKey"`
	Precio_Unitario float32 `gorm:"type:decimal(60,4);not null"`
	Cantidad        float32 `gorm:"type:decimal(60,4);not null"`
	Total           float32 `gorm:"type:decimal(60,4);not null"`
	Nombre          string  `gorm:"type:varchar(550);not null"`
	Id_Product      int     `gorm:"type:int(150);not null"`
	Id_Order        int     `gorm:"type:int(150);not null"`
}

type OrderDetails []OrderDetail
