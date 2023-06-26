package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// To generate UUID (Only use for postgresdb): CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type YoutubeTranscripts struct {
	Id          uint64         `gorm:"column:id;size:11;primary_key;auto_increment" json:"id"`
	UUID        uuid.UUID      `gorm:"type:string;default:(UUID());not null;index:idx_uuid" json:"uuid"`
	UserUUID    string         `gorm:"column:user_uuid;size:50;not null;index:idx_user_uuid" json:"user_uuid"`
	VideoId     string         `gorm:"column:videoId;size:50;not null;unique;index:idx_videoId" json:"video_id"`
	Transcripts []string       `gorm:"type:text;column:transcripts;not null" json:"transcripts"`
	Image       string         `gorm:"column:image;size:255" json:"image"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type CountYoutubeTranscripts struct {
	Count int64 `gorm:"column:count" json:"count"`
}
