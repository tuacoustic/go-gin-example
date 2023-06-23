package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersDto struct {
	gorm.Model
	Email     string    `form:"email" json:"email" xml:"email" binding:"required,gte=6,lte=255,email"`
	Phone     string    `form:"phone" json:"phone" xml:"phone" binding:"required,gte=6,lte=30"`
	Avatar    string    `form:"avatar" json:"avatar" xml:"avatar" binding:"required"`
	Password  string    `form:"password" json:"password" xml:"password" binding:"required,gte=6,lte=16"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUsersDto struct {
	gorm.Model
	Id    uint64 `form:"id"`
	Email string `form:"email" json:"email" xml:"email"`
	Phone string `form:"phone" json:"phone" xml:"phone"`
	Limit int    `form:"limit"` // Query Only
	Page  int    `form:"page"`  // Query Only
}

type UpdateUserDto struct {
	gorm.Model
	Email    string `form:"email" json:"email" xml:"email" binding:"omitempty,gte=6,lte=255,email"`
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"omitempty,gte=6,lte=30"`
	Avatar   string `form:"avatar" json:"avatar" xml:"avatar"`
	Password string `form:"password" json:"password" xml:"password" binding:"omitempty,gte=6,lte=16"`
}

func (u *UsersDto) BeforeSave(tx *gorm.DB) error {
	return hashPassword(&u.Password)
}

func (u *UpdateUserDto) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" {
		return hashPassword(&u.Password)
	}
	return nil
}

func hashPassword(password *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	*password = string(hashedPassword)
	return nil
}
