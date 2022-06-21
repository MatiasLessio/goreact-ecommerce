package productController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var productDto dto.ProductDto

	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, productsDto)
}

func GetProductsByIdCategory(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id_Category, _ := strconv.Atoi(c.Param("id"))
	var productsDto dto.ProductsDto

	productsDto, err := service.ProductService.GetProductsByIdCategory(id_Category)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}
