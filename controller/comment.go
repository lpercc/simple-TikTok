package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	repository.Response
	CommentList []repository.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	repository.Response
	Comment repository.Comment `json:"comment,omitempty"`
}

// CommentAction add comments
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := repository.UsersLoginInfo(token); exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
			comment := repository.Comment{
				Id:         repository.CommentActionAdd(text, user.Id, videoId),
				User:       user,
				Content:    text,
				CreateDate: time.Now().Format("01-02"),
			}
			c.JSON(http.StatusOK, CommentActionResponse{Response: repository.Response{StatusCode: 0},
				Comment: comment})
			return
		}
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have a comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    repository.Response{StatusCode: 0},
		CommentList: repository.FeedCommentList(videoId),
	})
}
