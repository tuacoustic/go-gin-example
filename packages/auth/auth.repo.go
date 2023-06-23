package auth

type AuthRepoIF interface {
	Login(AuthDto) (AuthResponseDto, error)
}
