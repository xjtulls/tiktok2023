package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

//var userIdSequence = int64(1)

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	userService := service.UserService{}
	resp, _ := userService.HandleRegister(c)
	c.JSON(http.StatusOK, resp)
}

func Login(c *gin.Context) {
	userService := service.UserService{}
	resp, _ := userService.HandleLogin(c)
	c.JSON(http.StatusOK, resp)
}

func UserInfo(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleGetUser(c)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
