package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"github.com/lpercc/simple-TikTok/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)
	repository.ConnectAndCheck()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
