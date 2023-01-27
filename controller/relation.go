package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"net/http"
)

type UserListResponse struct {
	repository.Response
	UserList []repository.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{repository.DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{repository.DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: repository.Response{
			StatusCode: 0,
		},
		UserList: []repository.User{repository.DemoUser},
	})
}
