package ytbtransconstants

import (
	"fmt"
)

func InputCrypto(videoId string, timestamp int64) string {
	// unixTime := time.Now().Unix()
	return fmt.Sprintf("video_id=%s&timestamp=%d", videoId, timestamp)
}

func GetTranscriptsUrl(url string, videoId string, timestamp int64, token string) string {
	return fmt.Sprintf("%s/api/youtube/get-transcript?video_id=%s&timestamp=%d&token=%s", url, videoId, timestamp, token)
}
