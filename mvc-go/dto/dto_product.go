package dto

type ProductDto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"base_price"`
	Id_Category int     `json:"id_category"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

type ProductsDto []ProductDto
