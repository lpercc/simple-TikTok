package repository

import (
	"gorm.io/gorm"
	"time"
)

type Videolists struct {
	Id            int64 `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

func FeedVedioList(Maxnum int64) (Videos []Video) {
	var video Video
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
		video.Id = videoDb.Id
		video.Author = user
		video.PlayUrl = videoDb.PlayUrl
		video.CoverUrl = videoDb.CoverUrl
		video.FavoriteCount = videoDb.FavoriteCount
		video.CommentCount = videoDb.CommentCount
		video.IsFavorite = videoDb.IsFavorite
		//fmt.Println(video)
		Videos = append(Videos, video)
	}
	return Videos
}
