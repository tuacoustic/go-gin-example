package users

import "gorm.io/gorm"

type UsersDto struct {
	gorm.Model
	Email    string `form:"email" json:"email" xml:"email" binding:"required,lte=255,email"`
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required,gte=6,lte=30"`
	Avatar   string `form:"avatar" json:"avatar" xml:"avatar" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required,len=gte=6,lte=16"`
}

type GetUsersDto struct {
	gorm.Model
	Email string `form:"email" json:"email" xml:"email" binding:"required,gte=6,lte=100"`
	Phone string `form:"phone" json:"phone" xml:"phone" binding:"required,gte=6,lte=30"`
}
