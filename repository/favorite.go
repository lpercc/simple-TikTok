package repository

import "gorm.io/gorm"

type FavoriteList struct {
	gorm.Model
	UserId  int64
	VideoId int64
}

// FavoriteActionAdd 点赞操作
func FavoriteActionAdd(userId int64, videoId int64) {
	//添加favoritelists 记录
	if GetDB().Create(&FavoriteList{UserId: userId, VideoId: videoId}).Error != nil {
		panic("failed to insert favorite info")
	}
	//video FavoriteCount+1
	if GetDB().Model(&VideoList{}).Where("id=?", videoId).Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error != nil {
		panic("failed to update Videolist table")
	}
}

// FavoriteActionCancel 取消点赞
func FavoriteActionCancel(userId int64, videoId int64) {
	//删除 favoritelists 记录
	var f FavoriteList
	if GetDB().Where("user_id=?", userId).Where("video_id=?", videoId).Delete(&f).Error != nil {
		panic("failed to insert favorite info")
	}
	//video FavoriteCount-1
	if GetDB().Model(&VideoList{}).Where("id=?", videoId).Update("favorite_count", gorm.Expr("favorite_count-?", 1)).Error != nil {
		panic("failed to update Videolist table")
	}
}

// feed favorite video Lists
func FeedFavoriteLists(userid int64) (Videos []Video) {
	var favoriteInf []FavoriteList
	//由userid查找favorite信息
	if GetDB().Find(&favoriteInf, "user_id=?", userid).Error != nil {
		panic("Author can't be find")
	}
	//遍历favoriteInf，格式化填入Videos
	for _, f := range favoriteInf {
		Videos = append(Videos, FeedVideo(f.VideoId, userid))
	}
	return Videos
}
