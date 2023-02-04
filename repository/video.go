package repository

import (
	"log"

	"gorm.io/gorm"
)

type VideoList struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

func SaveVideo(newVideo *VideoList) {
	err := GetDB().Create(newVideo)
	if err != nil {
		log.Println("Insert user error", err)
	}
}

func FeedVedioList(userId int64) (Videos []Video) {
	var count int64
	Maxnum := int64(30)
	// video record count
	GetDB().Model(&VideoList{}).Count(&count)
	if Maxnum > count {
		Maxnum = count
	}
	for i := int64(0); i < Maxnum; i++ {
		Videos = append(Videos, FeedVideo(count-i, userId))
	}
	return Videos
}

// FeedVideo feed only one video Inf
func FeedVideo(videoId int64, userId int64) (video Video) {
	var (
		videoDb VideoList
		author  User
		count   int64
	)
	// search a video record by videoid
	if err := GetDB().Find(&videoDb, videoId).Error; err != nil {
		return video
	}
	// search a user record by authorid
	if GetDB().Find(&author, videoDb.AuthorId).Error != nil {
		panic("Author can't be find")
	}
	// search a favorite record by userid and videoid
	if GetDB().Model(&FavoriteList{}).Where("user_id=?", userId).Where("video_id=?", videoId).Count(&count).Error != nil {
		panic("failed to find a favorite record")
	}
	// the video is user's favorite,user is existed
	if count == 1 && userId != -1 {
		videoDb.IsFavorite = true
	}
	// format to video interface
	return Video{
		Id:     int64(videoDb.ID),
		Author: author,
		//数据库中videoDb.PlayUrl是相对地址，video.PlayUrl需要带本机IP和端口的绝对地址，
		//视频是在本地Public文件夹
		// 1内网访问
		//PlayUrl:       "http://" + LOCALIPV_4 + ":8080/" + videoDb.PlayUrl,
		//CoverUrl:      "http://" + LOCALIPV_4 + ":8080/" + videoDb.CoverUrl,
		// 2远程访问
		PlayUrl:       "http://1af0010a.r6.cpolar.top/" + videoDb.PlayUrl,
		CoverUrl:      "http://1af0010a.r6.cpolar.top/" + videoDb.CoverUrl,
		FavoriteCount: videoDb.FavoriteCount,
		CommentCount:  videoDb.CommentCount,
		IsFavorite:    videoDb.IsFavorite,
	}
}
