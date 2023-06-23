package auth

import (
	"fmt"

	"gorm.io/gorm"
)

type repoAuthCRUD struct {
	db *gorm.DB
}

func AuthRepo(db *gorm.DB) *repoAuthCRUD {
	return &repoAuthCRUD{db}
}

func (repo *repoAuthCRUD) Login(userInput AuthDto) (AuthResponseDto, error) {
	fmt.Println("Demo", userInput)
	return AuthResponseDto{}, nil
}
