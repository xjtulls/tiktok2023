package service

import (
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

type RegisterAndLoginResopnse struct {
	Response
	UserId uint64 `json:"userId"`
	Token  string `json:"token"`
}

type GetUserResponse struct {
	Response
	model.User
}

func (u UserService) HandleRegister(c *gin.Context) (resp *RegisterAndLoginResopnse, err error) {
	username := c.Query("username")
	password := c.Query("password")
	// TODO 需要一个方法来获取token
	token := username + password
	// TODO 需要让数据库里面不能存储明文密码
	user := model.TableUser{
		Name:            username,
		Avatar:          AvaterDefault,
		Signature:       SignatureDefault,
		BackgroundImage: BackgroundImageDefault,
		Password:        password,
	}
	// 当数据中没有这个人的时候，会返回错误
	if err := model.Db.Where("name = ?", username).First(&user).Error; err == nil {
		return &RegisterAndLoginResopnse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  model.NotExitsEoor,
			},
			UserId: 0,
			Token:  "",
		}, err
	}
	// 到这里时用户不存在，即一个新用户
	if err := model.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &RegisterAndLoginResopnse{
		Response: Response{
			StatusCode: 0,
			//StatusMsg:  "",
		},
		UserId: user.ID,
		Token:  token,
	}, nil
}

func (u UserService) HandleLogin(c *gin.Context) (resp *RegisterAndLoginResopnse, err error) {
	username := c.Query("username")
	password := c.Query("password")
	// TODO token
	token := username + password
	user := model.TableUser{}
	// 当数据中没有这个人的时候，会返回错误
	if err := model.Db.Where("name = ?", username).First(&user).Error; err != nil {
		return &RegisterAndLoginResopnse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  model.NotExitsEoor,
			},
			UserId: 0,
			Token:  "",
		}, err
	}
	// 到这里时用户存在，判断密码是否匹配
	if password != user.Password {
		return &RegisterAndLoginResopnse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "密码错误",
			},
		}, err
	}
	return &RegisterAndLoginResopnse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserId: user.ID,
		Token:  token,
	}, nil
}

func HandleGetUser(c *gin.Context) {

}
