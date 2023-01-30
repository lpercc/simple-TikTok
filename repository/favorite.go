package repository

import "gorm.io/gorm"

type Favoritelists struct {
	gorm.Model
	UserId  int64
	VideoId int64
}

// FavoriteActionAdd 点赞操作
func FavoriteActionAdd(userId int64, videoId int64) {
	//添加favoritelists 记录
	if db.Create(&Favoritelists{UserId: userId, VideoId: videoId}).Error != nil {
		panic("failed to insert favorite info")
	}
	//video FavoriteCount+1
	if db.Model(&Videolists{}).Where("id=?", videoId).Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error != nil {
		panic("failed to update VideoLists table")
	}
}

// FavoriteActionCancel 取消点赞
func FavoriteActionCancel(userId int64, videoId int64) {
	//删除 favoritelists 记录
	if db.Delete(&Favoritelists{UserId: userId, VideoId: videoId}).Error != nil {
		panic("failed to insert favorite info")
	}
	//video FavoriteCount-1
	if db.Model(&Videolists{}).Where("id=?", videoId).Update("favorite_count", gorm.Expr("favorite_count-?", 1)).Error != nil {
		panic("failed to update VideoLists table")
	}
}

// feed favorite video Lists
func FeedFavoriteLists(userid int64) (Videos []Video) {
	var favoriteInf []Favoritelists
	var user User
	//获取用户信息
	if db.Find(&user, "user_id=?", userid).Error != nil {
		panic("Author can't be find")
	}
	//由userid查找favorite信息
	if db.Find(&favoriteInf, "user_id=?", userid).Error != nil {
		panic("Author can't be find")
	}
	//遍历favoriteInf，格式化填入Videos
	for _, f := range favoriteInf {
		var videoDb Videolists
		//获取对应视频信息，若找不到返回空切片
		if err := db.Find(&videoDb, "id=?", f.VideoId).Error; err != nil {
			return Videos
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
