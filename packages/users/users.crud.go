package users

import (
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/channel"
	"github.com/tuacoustic/go-gin-example/utils/console"
	"gorm.io/gorm"
)

type repoUsersCRUD struct {
	db *gorm.DB
}

func UsersRepo(db *gorm.DB) *repoUsersCRUD {
	return &repoUsersCRUD{db}
}

func (repo *repoUsersCRUD) Create(userInput UsersDto) (UsersDto, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = repo.db.Debug().Table("users").Create(&userInput).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		return userInput, nil
	}
	return UsersDto{}, err
}

func (repo *repoUsersCRUD) GetAll(queryParams interface{}) (entities.User, error) {
	var err error
	console.Info(queryParams)
	// done := make(chan bool)
	// go func(ch chan<- bool) {
	// 	defer close(ch)
	// 	err = repo.db.Debug().Table("users").Get().Error
	// 	if err != nil {
	// 		ch <- false
	// 		return
	// 	}
	// 	ch <- true
	// }(done)
	// if channel.OK(done) {
	// 	return userInput, nil
	// }
	return entities.User{}, err
}
