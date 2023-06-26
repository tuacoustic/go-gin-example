package youtubetranscripts

import "github.com/tuacoustic/go-gin-example/entities"

type YoutubeTransriptsRepoIF interface {
	Create(YoutubeTranscriptsDto, entities.User) (GetYoutubeTranscriptsDto, error)
	GetDetail(YtbTransQueryParamsDto) (GetYoutubeTranscriptsDto, error)
}
