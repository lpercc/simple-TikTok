package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"net/http"
	"time"
)

type FeedResponse struct {
	repository.Response
	VideoList []repository.Video `json:"video_list,omitempty"`
	NextTime  int64              `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	var userId int64
	// if user is not exist,userId = -1
	if user, exist := usersLoginInfo[token]; exist {
		userId = user.Id
	} else {
		userId = -1
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  repository.Response{StatusCode: 0},
		VideoList: repository.FeedVedioList(userId),
		NextTime:  time.Now().Unix(),
	})
}
