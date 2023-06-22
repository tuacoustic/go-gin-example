package users

import (
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/channel"
	tablename "github.com/tuacoustic/go-gin-example/utils/constants/tableName"
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
		err = repo.db.Debug().Table(tablename.TableName().Users).Create(&userInput).Error
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

func (repo *repoUsersCRUD) GetAll(queryParams GetUsersDto) ([]entities.User, error) {
	var err error
	var usersData []entities.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = repo.db.Debug().Table(tablename.TableName().Users).Find(&usersData).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		return usersData, nil
	}
	return []entities.User{}, err
}
