package services


import (
	addressCliente "mvc-go/clients/address" 
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"
)

type addressService struct{}

type addressServiceInterface interface {
	GetAddressById(id int) (dto.AddressDto, e.ApiError)
	GetAddresses() (dto.AddressesDto, e.ApiError)
}

var (
	AddressService addressServiceInterface
)

func init() {
	AddressService = &addressService{}
}

func (s *addressService) GetAddressById(id int) (dto.AddressDto, e.ApiError) {

	var address model.Address = addressCliente.GetAddressById(id) 
	var addressDto dto.AddressDto

	if address.Id == 0 {
		return addressDto, e.NewBadRequestApiError("address not found")
	}
	addressDto.Id = address.Id
	addressDto.Number = address.Number
	addressDto.Neighborhood = address.Neighborhood
	addressDto.Street = address.Street
	addressDto.Id_User = address.Id_User
	return addressDto, nil
}

func (s *addressService) GetAddresses() (dto.AddressesDto, e.ApiError) {

	var addresses model.Addresses = addressCliente.GetAddresses()
	var addressesDto dto.AddressesDto

	for _, address := range addresses {
		var addressDto dto.AddressDto
		addressDto.Id = address.Id
		addressDto.Number = address.Number
		addressDto.Neighborhood = address.Neighborhood
		addressDto.Street = address.Street

		addressDto.Id_User = address.Id_User
		addressesDto = append(addressesDto, addressDto)
	}

	return addressesDto, nil
}