package dto

type OrderDetailDto struct {
	Id              int     `json:"id"`
	Precio_Unitario float32 `json:"precio_unitario"`
	Cantidad        float32 `json:"cantidad"`
	Total           float32 `json:"total"`
	Nombre          string  `json:"nombre"`
	Id_Producto     int     `json:"id_product"`
	Id_Order        int     `json:"id_order"`
}

type OrderDetailsDto []OrderDetailDto
