package youtubetranscripts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tuacoustic/go-gin-example/entities"
	"github.com/tuacoustic/go-gin-example/utils/channel"
	tablename "github.com/tuacoustic/go-gin-example/utils/constants/tableName"
	ytbtransconstants "github.com/tuacoustic/go-gin-example/utils/constants/ytbTransConstants"
	cryptos "github.com/tuacoustic/go-gin-example/utils/cryptoS"
	"github.com/tuacoustic/go-gin-example/utils/helper"
	"github.com/tuacoustic/go-gin-example/utils/setting"
	"gorm.io/gorm"
)

type repoYoutubeTranscriptsCRUD struct {
	db *gorm.DB
}

func YoutubeTransriptsRepo(db *gorm.DB) *repoYoutubeTranscriptsCRUD {
	return &repoYoutubeTranscriptsCRUD{db}
}

func (repo *repoYoutubeTranscriptsCRUD) Create(userInput YoutubeTranscriptsDto, UserInfo entities.User) (GetYoutubeTranscriptsDto, error) {
	var err error
	var ytbTrans entities.YoutubeTranscripts
	var transcript []TranscriptDto
	unixTime := time.Now().Unix()
	// Get Transcript

	// From Python
	inputCrypto := ytbtransconstants.InputCrypto(userInput.VideoId, unixTime)
	token := cryptos.Generate(inputCrypto, setting.AppSetting.CryptoKey)
	url := ytbtransconstants.GetTranscriptsUrl(setting.AppSetting.PythonUrl, userInput.VideoId, unixTime, token)
	respBody, err := helper.GetRequest(url)
	if err != nil {
		return GetYoutubeTranscriptsDto{}, err
	}

	// Unmarshal the Json data into the Transcript variable
	err = json.Unmarshal(respBody, &transcript)
	if err != nil {
		return GetYoutubeTranscriptsDto{}, err
	}
	ytbTrans = entities.YoutubeTranscripts{
		UserUUID:    fmt.Sprintf("%s", UserInfo.UUID),
		Title:       userInput.Title,
		Desc:        userInput.Desc,
		VideoId:     userInput.VideoId,
		Transcripts: respBody,
	}

	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if err = repo.db.Debug().Table(tablename.TableName().YoutubeTranscripts).Create(&ytbTrans).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		resp := GetYoutubeTranscriptsDto{
			UserUUID:    ytbTrans.UserUUID,
			Title:       ytbTrans.Title,
			Desc:        ytbTrans.Desc,
			VideoId:     ytbTrans.VideoId,
			Transcripts: transcript,
		}
		return resp, nil
	}
	return GetYoutubeTranscriptsDto{}, err
}

func (repo *repoYoutubeTranscriptsCRUD) GetDetail(queryPrams YtbTransQueryParamsDto) (GetYoutubeTranscriptsDto, error) {
	var err error
	var ytbTransData entities.YoutubeTranscripts
	var transcript []TranscriptDto
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		if err = repo.db.Debug().Table(tablename.TableName().YoutubeTranscripts).Where("uuid = ?", queryPrams.UUID).First(&ytbTransData).Error; err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channel.OK(done) {
		// Unmarshal the Json data into the Transcript variable
		err = json.Unmarshal(ytbTransData.Transcripts, &transcript)
		if err != nil {
			return GetYoutubeTranscriptsDto{}, err
		}
		resp := GetYoutubeTranscriptsDto{
			Id:          ytbTransData.Id,
			UUID:        ytbTransData.UUID,
			UserUUID:    ytbTransData.UserUUID,
			VideoId:     ytbTransData.VideoId,
			Title:       ytbTransData.Title,
			Desc:        ytbTransData.Desc,
			Transcripts: transcript,
			Image:       ytbTransData.Image,
			DeletedAt:   ytbTransData.DeletedAt,
			CreatedAt:   ytbTransData.CreatedAt,
			UpdatedAt:   ytbTransData.UpdatedAt,
		}
		return resp, nil
	}
	return GetYoutubeTranscriptsDto{}, err
}
