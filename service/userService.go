package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

const (
	AvaterDefault = ""
	SignatureDefault
	BackgroundImageDefault
)

type UserService struct {
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type RegisterResopnse struct {
	Response
	UserID uint64 `json:"userId"`
	Token  string `json:"token"`
}

func (u UserService) HandleRegister(c *gin.Context) (resp *RegisterResopnse, err error) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	user := model.User{
		Name:            username,
		Avatar:          AvaterDefault,
		Signature:       SignatureDefault,
		BackgroundImage: BackgroundImageDefault,
		Password:        password,
	}
	// 当数据中没有这个人的时候，会返回错误
	if err := model.Db.Where("name = ?", username).First(&user).Error; err == nil {
		return &RegisterResopnse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("%v", err),
			},
			UserID: -1,
			Token:  "",
		}, err
	}
	// 到这里时用户不存在，即一个新用户
	if err := model.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &RegisterResopnse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserID: user.ID,
		Token:  token,
	}, nil
}
