package main

import "github.com/lpercc/simple-TikTok/repository"

func main() {
	//go service.RunMessageServer()

	//r := gin.Default()

	//initRouter(r)
	//repository.Initgorm()
	repository.ConnectAndCheck()
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
