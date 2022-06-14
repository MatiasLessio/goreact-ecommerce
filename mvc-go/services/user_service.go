package services

import (
	userClient "mvc-go/clients/user"
	"mvc-go/dto"
	"mvc-go/model"
	e "mvc-go/utils/errors"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}


var jwtKey = []byte("secret_key")
func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {

	log.Debug(loginDto)
	var user model.User = userClient.GetUserByUserName(loginDto.User, loginDto.Password)
	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	if user.Password == loginDto.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
	}
	return tokenDto, nil
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}

	userDto.Name = user.Name
	userDto.LastName = user.Last_Name
	userDto.UserName = user.User_Name
	userDto.Password = user.Password
	userDto.Id = user.Id
	return userDto, nil
}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto
		userDto.Name = user.Name
		userDto.LastName = user.Last_Name
		userDto.UserName = user.User_Name
		userDto.Password = user.Password
		userDto.Id = user.Id

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}
