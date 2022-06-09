package product

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order

	Db.Where("id = ?", id).First(&order) 
	log.Debug("Order: ", order)

	return order
}

func GetOrdersByIdUser(id_User int) model.Orders {
	var orders model.Orders

	log.Debug("id_User: ", id_User)
	Db.Where("id_user = ?", id_User).Find(&orders)
	log.Debug("Order: ", orders)

	return orders
}

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)

	log.Debug("Orders: ", orders)

	return orders
}
func UpdateMontoFinal(monto float32, id_Order int) float32 {
	result := Db.Model(&model.Order{}).Where("id = ?", id_Order).Update("monto_final", monto)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("Order no encontrada")
	}
	return monto
}

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Order Created: ", order.Id)
	return order
}
