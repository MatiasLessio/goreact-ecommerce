package services

import (
	productCliente "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetProductsByIdCategory(id_Category int) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id)
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.Name = product.Name
	productDto.Price = product.Price
	productDto.Id = product.Id
	productDto.Description = product.Description
	productDto.Stock = product.Stock
	productDto.Id_Category = product.Id_Category
	return productDto, nil
}

//muestra solo los productos de la categoria solicitada
func (s *productService) GetProductsByIdCategory(id_Category int) (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProductsByIdCategory(id_Category)
	var productsDto dto.ProductsDto
	
	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("products not found")
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.Id_Category = product.Id_Category

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

//muestra todos los productos
func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	if len(products) == 0 {
		return productsDto, e.NewBadRequestApiError("products not found")
	}

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Description = product.Description
		productDto.Stock = product.Stock
		productDto.Id_Category = product.Id_Category

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}
