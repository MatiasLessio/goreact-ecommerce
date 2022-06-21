package orderController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderById(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var orderDto dto.OrderDto

	orderDto, err := service.OrderService.GetOrderById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, orderDto)
}

func GetOrders(c *gin.Context) {
	var ordersDto dto.OrdersDto
	ordersDto, err := service.OrderService.GetOrders()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, ordersDto)
}

func OrderInsert(c *gin.Context) {
	var orderDto dto.OrderDto

	err := c.BindJSON(&orderDto)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDto, er := service.OrderService.InsertOrder(orderDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDto)
}

func GetOrdersByIdUser(c *gin.Context) {
	log.Debug("Order id to load: " + c.Param("id"))

	id_User, _ := strconv.Atoi(c.Param("id"))
	var ordersDto dto.OrdersDto

	ordersDto, err := service.OrderService.GetOrdersByIdUser(id_User)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, ordersDto)
}
