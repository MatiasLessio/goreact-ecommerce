package app

import (
	addressController "mvc-go/controllers/address"
	categoryController "mvc-go/controllers/category"
	orderController "mvc-go/controllers/order"
	orderDetailController "mvc-go/controllers/order_detail"
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/login", userController.LoginUser) //hay que ver como lo hacemos bien

	//Address Mapping
	router.GET("/address/:id", addressController.GetAddressById)
	router.GET("/address", addressController.GetAddresses)

	//Category Mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/categories", categoryController.GetCategories)

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/orders", orderController.GetOrders)
	router.POST("/neworder", orderController.OrderInsert)
	router.GET("/orderUser/:idUser", orderController.GetOrdersByIdUser)

	//OrderDetail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/product/category/:id", productController.GetProductsByIdCategory)

	log.Info("Finishing mappings configurations")
}
