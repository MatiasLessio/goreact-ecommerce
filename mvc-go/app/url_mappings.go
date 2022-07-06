package app

import (
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
	router.GET("/users", userController.GetUsers)
	router.POST("/login", userController.LoginUser) //login del user

	//Category Mapping
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/categories", categoryController.GetCategories) //muestra todas las categorias

	//Order Mapping
	router.POST("/neworder", orderController.OrderInsert) //insert order
	router.GET("/order/user/:token", orderController.GetOrdersByIdUser)

	//OrderDetail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.GET("/orderDetails", orderDetailController.GetOrderDetails)

	//Product Mapping
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)                         //muestra todos los productos
	router.GET("/product/category/:id", productController.GetProductsByIdCategory) //muestra los productos de una categoria

	log.Info("Finishing mappings configurations")
}

//'<script>alert("XSS")</script>'
