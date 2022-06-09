package dto

type CategoryDto struct {
	Id_Category int    `json:"id_category"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type CategoriesDto []CategoryDto
