package orderDetail

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderDetailById(id int) model.OrderDetail {
	var orderDetail model.OrderDetail

	Db.Where("id = ?", id).First(&orderDetail)
	log.Debug("OrderDetail: ", orderDetail)

	return orderDetail
}

func GetOrderDetails() model.OrderDetails {
	var orderDetails model.OrderDetails
	Db.Find(&orderDetails)

	log.Debug("OrderDetails: ", orderDetails)

	return orderDetails
}

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("OrderDetail Created: ", orderDetail.Id)
	result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.Id_Product).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))

	if result1.Error != nil {
		log.Error("Producto no encontrado")
	}
	return orderDetail
}

func GetOrderDetailByIdOrder(id_Order int) model.OrderDetails {
	var ordersDetail model.OrderDetails

	Db.Where("id_order = ?", id_Order).Find(&ordersDetail)
	log.Debug("OrderDetail: ", ordersDetail)

	return ordersDetail
}
func InsertOrdersDetail(ordersDetail model.OrderDetails) model.OrderDetails {

	for _, orderDetail := range ordersDetail {

		result := Db.Create(&orderDetail)

		if result.Error != nil {
			log.Error("Error al crear")
		}

		log.Debug("Order_Detail Created: ", orderDetail.Id)
		result1 := Db.Model(&model.Product{}).Where("id = ?", orderDetail.Id_Product).Update("stock", gorm.Expr("stock - ? ", orderDetail.Cantidad))
		if result1.Error != nil {
			log.Error("Producto no encontrado")
		}
	}

	return ordersDetail
}
