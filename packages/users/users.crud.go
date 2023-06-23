package users

import (
	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/repositories"
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
		if err = repo.db.Debug().Table(tablename.TableName().Users).Create(&userInput).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		userInput.Password = ""
		return userInput, nil
	}
	return UsersDto{}, err
}

func (repo *repoUsersCRUD) GetAll(queryParams GetUsersDto) ([]entities.User, int, error) {
	var err error
	var usersData []entities.User
	var count int64

	// Query from where
	queryUser := GetUsersDto{
		Id: queryParams.Id,
	}

	// Query from like
	queryUserLike := GetUsersDto{
		Email: "%" + queryParams.Email + "%",
		Phone: "%" + queryParams.Phone + "%",
	}

	// Query
	query := repo.db.Table(tablename.TableName().Users).
		Where("email LIKE ? and phone LIKE ?", queryUserLike.Email, queryUserLike.Phone).
		Where(&queryUser).
		Order("created_at desc")

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		// Execute the count query
		if err = query.Debug().Count(&count).Error; err != nil {
			ch <- false
			return
		}
		paginatedQuery, getErr := repositories.Paginate(query, queryParams.Limit, queryParams.Page)
		if getErr != nil {
			ch <- false
			return
		}
		// Execute the paginated query
		if err = paginatedQuery.Debug().Find(&usersData).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		return usersData, int(count), nil
	}
	return []entities.User{}, 0, err
}

func (repo *repoUsersCRUD) Update(userId string, userInput UpdateUserDto) (entities.User, error) {
	var err error
	var userData entities.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if err = repo.db.Debug().Table(tablename.TableName().Users).Where("id = ?", userId).Updates(&userInput).Error; err != nil {
			ch <- false
			return
		}
		if err = repo.db.Debug().Table(tablename.TableName().Users).Where("id = ?", userId).First(&userData).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		return userData, nil
	}
	return entities.User{}, err
}

func (repo *repoUsersCRUD) SoftDelete(userId string) (bool, error) {
	var err error
	var user entities.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if err = repo.db.Debug().Table(tablename.TableName().Users).Where("id = ?", userId).First(&user).Error; err != nil {
			ch <- false
			return
		}
		if err = repo.db.Delete(&user).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		return true, nil
	}
	return false, err
}
