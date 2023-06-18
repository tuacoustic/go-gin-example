package users

import "github.com/tuacoustic/go-gin-example/entities"

type UsersRepoIF interface {
	Create(UsersDto) (UsersDto, error)
	GetAll(interface{}) (entities.User, error)
}
