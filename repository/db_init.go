package repository

import (
	"fmt"
	"log"
	"net"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var LOCALIPV_4 string
var db *gorm.DB

func ConnectAndCheck() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/simple-tiktok?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})
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
func GetLocalIP() (ipv4 string, err error) {
	// var (
	// 	addrs []net.Addr
	// 	addr net.Addr
	// 	ipNet *net.IPNet // IP地址
	// 	isIpNet bool
	// )
	// 获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr := range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		ipNet, isIpNet := addr.(*net.IPNet)
		if ipNet.IP.IsLoopback() {
			continue
		}
		if !isIpNet {
			continue
		}
		// 跳过IPV6
		if ipNet.IP.To4() == nil {
			continue
		} else {
			ipv4 = ipNet.IP.String()
		}
		if strings.HasPrefix(ipv4, "192.168.") {
			// 192.168.1.1
			fmt.Println("Local IPv4:", ipv4)
			return
		}
	}
	return
}
