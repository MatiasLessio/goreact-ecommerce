package product

import (
	"mvc-go/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var product model.Product

	Db.Where("id = ?", id).First(&product)
	log.Debug("Product: ", product)

	return product
}

func GetProductsByIdCategory(id_Category int) model.Products { //SELECT * from products WHERE id_category = x;
	var products model.Products
	log.Debug("id_category: ", id_Category)
	Db.Where("id_category = ?", id_Category).Find(&products)
	log.Debug("Product: ", products)
	return products
}


func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)
	log.Debug("Products: ", products)

	return products
}
