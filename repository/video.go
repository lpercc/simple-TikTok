package repository

import (
	"gorm.io/gorm"
)

type Videolists struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

func FeedVedioList(Maxnum int64) (Videos []Video) {
	//var video Video
	var count int64
	db.Model(&Videolists{}).Count(&count)
	if Maxnum > count {
		Maxnum = count
	}
	for i := int64(0); i < Maxnum; i++ {
		var (
			videoDb Videolists
			user    User
		)
		if err := db.Find(&videoDb, count-i).Error; err != nil {
			return Videos
		}
		if db.Find(&user, videoDb.AuthorId).Error != nil {
			panic("Author can't be find")
		}
		//fmt.Println(i, videoDb)
		video := Video{
			Id:     int64(videoDb.ID),
			Author: user,
			//数据库中videoDb.PlayUrl是相对地址，video.PlayUrl需要带本机IP和端口的绝对地址，
			//视频是在本地Public文件夹
			PlayUrl:       "http://" + LocalIp + ":8080/" + videoDb.PlayUrl,
			CoverUrl:      "http://" + LocalIp + ":8080/" + videoDb.CoverUrl,
			FavoriteCount: videoDb.FavoriteCount,
			CommentCount:  videoDb.CommentCount,
			IsFavorite:    videoDb.IsFavorite,
		}
		//fmt.Println(video)
		Videos = append(Videos, video)
	}
	return Videos
}
