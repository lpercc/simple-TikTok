package repository

import (
	"github.com/lpercc/simple-TikTok/controller"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	videoInf controller.User
}
