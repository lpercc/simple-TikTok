package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lpercc/simple-TikTok/repository"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=TestUser, password=123456

type UserLoginResponse struct {
	repository.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	repository.Response
	User repository.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := repository.UsersLoginInfo(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := repository.User{
			Name:  username,
			Token: token,
		}
		repository.AddUser(newUser)
		user, _ := repository.UsersLoginInfo(token)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := repository.UsersLoginInfo(token); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := repository.UsersLoginInfo(token); exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
