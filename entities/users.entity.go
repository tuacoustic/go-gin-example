package entities

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Street  string
	City    string
	Country string
}

// To generate UUID (Only use for postgresdb): CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type User struct {
	Id uint64 `gorm:"column:id;size:11;primary_key;auto_increment" json:"id"`
	// UUID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();not null" json:"uuid"` Only use for postgresdb
	UUID      uuid.UUID `gorm:"type:string;default:(UUID());not null" json:"uuid"`
	Email     string    `gorm:"column:email;size:255;not null;unique;index:idx_email" json:"email"`
	Phone     string    `gorm:"column:phone;size:30;not null;unique;index:idx_phone" json:"phone"`
	Avatar    string    `gorm:"column:avatar;size:255" json:"avatar"`
	Password  string    `gorm:"column:password;size:100;not null"`
	Address   Address   `gorm:"embedded;embeddedPrefix:address_;comment: 'street: đường | city: thành phố | country: quốc gia'"`
	DeletedAt time.Time `json:"deleted_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// type:timestamp; default:current_timestamp()
}

type CountUsers struct {
	Count int64 `gorm:"column:count" json:"count"`
}
