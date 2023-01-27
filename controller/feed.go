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
	c.JSON(http.StatusOK, FeedResponse{
		Response:  repository.Response{StatusCode: 0},
		VideoList: repository.FeedVedioList(30),
		NextTime:  time.Now().Unix(),
	})
}
