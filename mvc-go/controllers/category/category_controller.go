package categoryController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCategoryById(c *gin.Context) {
	log.Debug("Category id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var categoryDto dto.CategoryDto

	categoryDto, err := service.CategoryService.GetCategoryById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, categoryDto)
}

func GetCategories(c *gin.Context) {
	var categoriesDto dto.CategoriesDto
	categoriesDto, err := service.CategoryService.GetCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, categoriesDto)
}
