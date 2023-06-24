package auth

import (
	"errors"

	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/middlewares"
	"github.com/tuacoustic/go-gin-example/utils/channel"
	tablename "github.com/tuacoustic/go-gin-example/utils/constants/tableName"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repoAuthCRUD struct {
	db *gorm.DB
}

func AuthRepo(db *gorm.DB) *repoAuthCRUD {
	return &repoAuthCRUD{db}
}

func (repo *repoAuthCRUD) Login(userInput AuthDto) (AuthResponseDto, error) {
	var err error
	var userData entities.User

	query := repo.db.Table(tablename.TableName().Users).
		Where("email LIKE ? OR phone LIKE ?", userInput.Input, userInput.Input)

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if err = query.Debug().First(&userData).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		// Compare password
		err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(userInput.Password))
		if err != nil {
			return AuthResponseDto{}, errors.New("invalid password")
		}

		// Generate Token
		token, err := middlewares.GenerateToken(userData.UUID)
		if err != nil {
			return AuthResponseDto{}, err
		}
		response := AuthResponseDto{
			Email:  userData.Email,
			Phone:  userData.Phone,
			Avatar: userData.Avatar,
			Token:  token,
		}
		return response, nil
	}
	return AuthResponseDto{}, nil
}

func (repo *repoAuthCRUD) Profile(token string) (entities.User, error) {
	return entities.User{}, nil
}
