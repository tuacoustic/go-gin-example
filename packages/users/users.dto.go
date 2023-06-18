package users

import "gorm.io/gorm"

type UsersDto struct {
	gorm.Model
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Avatar   string `form:"avatar" json:"avatar" xml:"avatar" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
