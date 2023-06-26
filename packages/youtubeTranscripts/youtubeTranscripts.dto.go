package youtubetranscripts

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type YoutubeTranscriptsDto struct {
	VideoId string `form:"video_id" json:"video_id" xml:"video_id" binding:"required"`
	Title   string `form:"title" json:"title" xml:"title" binding:"required,lte=255"`
	Desc    string `form:"desc" json:"desc" xml:"desc,lte=255"`
}

type TranscriptDto struct {
	Duration float32 `json:"duration"`
	Start    float32 `json:"start"`
	Text     string  `json:"text"`
}

type GetYoutubeTranscriptsDto struct {
	Id          uint64          `json:"id"`
	UUID        uuid.UUID       `form:"uuid" json:"uuid"`
	UserUUID    string          `json:"user_uuid"`
	VideoId     string          `json:"video_id"`
	Title       string          `json:"title"`
	Desc        string          `json:"desc"`
	Transcripts []TranscriptDto `json:"transcripts"`
	Image       string          `json:"image"`
	DeletedAt   gorm.DeletedAt  `json:"deletaed_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type YtbTransQueryParamsDto struct {
	UUID  string `form:"uuid"`
	Limit int    `form:"limit"` // Query Only
	Page  int    `form:"page"`  // Query Only
}
