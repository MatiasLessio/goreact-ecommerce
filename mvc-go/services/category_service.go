package services

import (
	categoryCliente "mvc-go/clients/category"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type categoryService struct{}

type categoryServiceInterface interface {
	GetCategoryById(id int) (dto.CategoryDto, e.ApiError)
	GetCategories() (dto.CategoriesDto, e.ApiError)
}

var (
	CategoryService categoryServiceInterface
)

func init() {
	CategoryService = &categoryService{}
}

func (s *categoryService) GetCategoryById(id int) (dto.CategoryDto, e.ApiError) {

	var category model.Category = categoryCliente.GetCategoryById(id)
	var categoryDto dto.CategoryDto

	if category.Id == 0 {
		return categoryDto, e.NewBadRequestApiError("category not found")
	}
	categoryDto.Descripcion = category.Description
	categoryDto.Nombre = category.Name
	categoryDto.Id_Category = category.Id
	return categoryDto, nil
}

func (s *categoryService) GetCategories() (dto.CategoriesDto, e.ApiError) {

	var categories model.Categories = categoryCliente.GetCategories()
	var categoriesDto dto.CategoriesDto

	for _, category := range categories {
		var categoryDto dto.CategoryDto
		categoryDto.Descripcion = category.Description
		categoryDto.Id_Category = category.Id
		categoryDto.Nombre = category.Name

		categoriesDto = append(categoriesDto, categoryDto)
	}

	return categoriesDto, nil
}
