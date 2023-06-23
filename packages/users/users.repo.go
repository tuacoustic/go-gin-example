package users

import "github.com/tuacoustic/go-gin-example/entities"

type UsersRepoIF interface {
	Create(UsersDto) (UsersDto, error)
	GetAll(GetUsersDto) ([]entities.User, int, error)
	Update(string, UpdateUserDto) (entities.User, error)
	SoftDelete(string) (bool, error)
}
