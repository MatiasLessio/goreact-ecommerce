package services

import (
	orderClient "mvc-go/clients/order"
	orderDetailClient "mvc-go/clients/order_detail"
	productClient "mvc-go/clients/product"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
	"time"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError)
	GetOrdersByIdUser(id_User int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {

	var order model.Order = orderClient.GetOrderById(id)
	var orderDto dto.OrderDto

	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	orderDto.Id = order.Id
	orderDto.Fecha = order.Fecha
	orderDto.Monto_Final = order.Monto_Final
	orderDto.Id_Usuario = order.Id_User

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderClient.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Fecha = order.Fecha
		orderDto.Monto_Final = order.Monto_Final
		orderDto.Id_Usuario = order.Id_User
		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) InsertOrder(orderDto dto.OrderDto) (dto.OrderDto, e.ApiError) {

	var order model.Order

	order.Monto_Final = orderDto.Monto_Final
	order.Fecha = time.Now()
	order.Id_User = orderDto.Id_Usuario

	order = orderClient.InsertOrder(order)

	var details model.OrderDetails
	var total float32

	for _, detailDto := range orderDto.Order_Details {

		var detail model.OrderDetail
		detail.Id_Product = detailDto.Id_Producto

		var product model.Product = productClient.GetProductById(detail.Id_Product)
		detail.Precio_Unitario = product.Price
		detail.Cantidad = detailDto.Cantidad
		detail.Total = detail.Precio_Unitario * detail.Cantidad
		detail.Nombre = product.Name
		detail.Id_Order = order.Id

		total = total + detail.Total

		details = append(details, detail)
	}

	orderClient.UpdateMontoFinal(total, order.Id)

	orderDetailClient.InsertOrdersDetail(details)

	return orderDto, nil
}

//get de orders por el id del cliente
func (s *orderService) GetOrdersByIdUser(id_User int) (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderClient.GetOrdersByIdUser(id_User)
	var ordersDto dto.OrdersDto
	if len(orders) == 0 {
		return ordersDto, e.NewBadRequestApiError("order not found")
	}
	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Fecha = order.Fecha
		orderDto.Monto_Final = order.Monto_Final
		orderDto.Id_Usuario = order.Id_User

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}
