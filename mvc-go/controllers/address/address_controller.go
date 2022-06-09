package addressController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAddressById(c *gin.Context) {
	log.Debug("Address id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var addressDto dto.AddressDto

	addressDto, err := service.AddressService.GetAddressById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, addressDto)
}

func GetAddresses(c *gin.Context) {
	var addressesDto dto.AddressesDto
	addressesDto, err := service.AddressService.GetAddresses()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, addressesDto)
}
