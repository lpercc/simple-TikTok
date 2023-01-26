package repository

import (
	"github.com/lpercc/simple-TikTok/controller"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	videoInf controller.Video
}

func FeedVedioList(Maxnum int64) (Videos []controller.Video) {
	var count int64
	db.Model(&Video{}).Count(&count)
	var video Video
	for i := int64(0); i < Maxnum; i++ {
		if db.Last(&video, count-i).Error != nil {
			return Videos
		}
		Videos = append(Videos, video.videoInf)
	}
	return Videos
}
