package auth

import "github.com/tuacoustic/go-gin-example/entities"

type AuthRepoIF interface {
	Login(AuthDto) (AuthResponseDto, error)
	Profile(string) (entities.User, error)
}
