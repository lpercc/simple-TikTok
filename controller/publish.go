package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"net/http"
	"path/filepath"
	"sync/atomic"
)

type VideoListResponse struct {
	repository.Response
	VideoList []repository.Video `json:"video_list"`
}

var videoIdSequence = int64(2) // video id sequence

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := repository.UsersLoginInfo(token); !exist {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user, _ := repository.UsersLoginInfo(token)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	atomic.AddInt64(&videoIdSequence, 1)
	newVideo := &repository.VideoList{
		AuthorId:      user.Id,
		PlayUrl:       "static/" + finalName,
		CoverUrl:      "static/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
	}

	repository.SaveVideo(newVideo)

	c.JSON(http.StatusOK, repository.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
