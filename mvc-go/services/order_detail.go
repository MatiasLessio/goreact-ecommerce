package services

import (
	orderDetailCliente "mvc-go/clients/order_detail"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetails() (dto.OrderDetailsDto, e.ApiError)
	InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetailByIdOrder(id_Order int) (dto.OrderDetailsDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

//get by id
func (s *orderDetailService) GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail = orderDetailCliente.GetOrderDetailById(id)
	var orderDetailDto dto.OrderDetailDto

	if orderDetail.Id == 0 {
		return orderDetailDto, e.NewBadRequestApiError("orderDetail not found")
	}
	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Cantidad = orderDetail.Cantidad
	orderDetailDto.Precio_Unitario = orderDetail.Precio_Unitario
	orderDetailDto.Total = orderDetail.Total
	orderDetailDto.Id_Order = orderDetail.Id_Order
	orderDetailDto.Id_Producto = orderDetail.Id_Product
	return orderDetailDto, nil
}

//get array
func (s *orderDetailService) GetOrderDetails() (dto.OrderDetailsDto, e.ApiError) {

	var orderDetails model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var orderDetailsDto dto.OrderDetailsDto

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Cantidad = orderDetail.Cantidad
		orderDetailDto.Precio_Unitario = orderDetail.Precio_Unitario
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.Id_Order = orderDetail.Id_Order
		orderDetailDto.Id_Producto = orderDetail.Id_Product

		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

//insert
func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail
	orderDetail.Cantidad = orderDetailDto.Cantidad
	orderDetail.Precio_Unitario = orderDetailDto.Precio_Unitario
	orderDetail.Total = orderDetailDto.Precio_Unitario * orderDetailDto.Cantidad
	orderDetail.Id_Order = orderDetailDto.Id_Order
	orderDetail.Id_Product = orderDetailDto.Id_Producto

	orderDetail = orderDetailCliente.InsertOrderDetail(orderDetail)

	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Total = orderDetail.Total

	return orderDetailDto, nil
}

//devuelve el detalle con el id de la order
func (s *orderDetailService) GetOrderDetailByIdOrder(id_Order int) (dto.OrderDetailsDto, e.ApiError) {

	var ordersDetail model.OrderDetails = orderDetailCliente.GetOrderDetailByIdOrder(id_Order)
	var ordersDetailDto dto.OrderDetailsDto

	for _, orderDetail := range ordersDetail {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Cantidad = orderDetail.Cantidad
		orderDetailDto.Precio_Unitario = orderDetail.Precio_Unitario
		orderDetailDto.Total = orderDetail.Total
		orderDetailDto.Id_Producto = orderDetail.Id_Product

		ordersDetailDto = append(ordersDetailDto, orderDetailDto)
	}
	return ordersDetailDto, nil
}
