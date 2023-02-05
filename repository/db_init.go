package repository

import (
	"log"
	"net"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var LOCALIPV_4 string
var db *gorm.DB

func ConnectAndCheck() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple-tiktok?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info),})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	if db.AutoMigrate(&VideoList{}, &User{}, &CommentList{}, &FavoriteList{}) != nil {
		panic("failed to create table")
	}
	_db, err := db.DB()
	if err != nil {
		log.Fatalln("db connected error ", err)
	}
	_db.SetMaxOpenConns(100)
	_db.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return db
}

// 获取本机网卡IP
func GetLocalIP() (string, error) {
    // 思路来自于Python版本的内网IP获取
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Println("internal IP fetch failed, detail:",err)
    }
    defer conn.Close()
 
    // udp 面向无连接，所以这些东西只在你本地捣鼓
    res := conn.LocalAddr().String()
    res = strings.Split(res, ":")[0]
	log.Println("internal IP fetch success, IP:",res)
    return res, nil
}
