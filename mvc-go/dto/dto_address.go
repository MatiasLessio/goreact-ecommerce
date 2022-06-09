package dto

type AddressDto struct {
	Id           int    `json:"id_address"`
	Street       string `json:"calle"`
	Number       int    `json:"numero"`
	Neighborhood string `json:"barrio"`
	Id_User      int    `json:"id_user"`
}
type AddressesDto []AddressDto
