package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var LocalIp = "192.168.101.17"
var db *gorm.DB

func ConnectAndCheck() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple-tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	if db.AutoMigrate(&Videolists{}, &User{}, &Comments{}, &Favoritelists{}) != nil {
		panic("failed to create table")
	}
	/*
		//Create
		var DemoVideos = []Videolists{
			{
				AuthorId:      1,
				PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
				CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			},
			{
				AuthorId:      1,
				PlayUrl:       "http://192.168.196.76:8080/static/2_wx_camera_1674529126546.mp4",
				CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
				FavoriteCount: 0,
				CommentCount:  0,
				IsFavorite:    false,
			},
		}
		db.Create(&DemoUser)
		db.Create(&DemoVideos)

	*/
}
