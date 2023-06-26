package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// To generate UUID (Only use for postgresdb): CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type YoutubeTranscripts struct {
	Id          uint64         `gorm:"column:id;size:11;primary_key;auto_increment" json:"id"`
	UUID        uuid.UUID      `gorm:"type:string;default:(UUID());not null" json:"uuid"`
	UserUUID    string         `gorm:"column:user_uuid;size:50;not null;index:idx_user_uuid" json:"user_uuid"`
	VideoId     string         `gorm:"column:videoId;size:50;not null;unique;index:idx_videoId" json:"video_id"`
	Title       string         `gorm:"column:title;size:255;not null;index:idx_title" json:"title"`
	Desc        string         `gorm:"column:desc;size:255;not null;index:idx_desc" json:"desc"`
	Transcripts []byte         `gorm:"type:text;column:transcripts;not null" json:"transcripts"`
	Image       string         `gorm:"column:image;size:255" json:"image"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Transcript struct {
	Duration float32 `json:"duration"`
	Start    float32 `json:"start"`
	Text     string  `json:"text"`
}
type CountYoutubeTranscripts struct {
	Count int64 `gorm:"column:count" json:"count"`
}
