package service

import (
	"swuxfdc/model"
	"swuxfdc/serializer"

	"github.com/gin-gonic/gin"
)

// UserChangeService 管理用户密码修改的服务
type UserChangeService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=40"`
	PasswordNew     string `form:"password_new" json:"password_new" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
}

// Change 用户密码修改函数
func (service *UserChangeService) Change(c *gin.Context) serializer.Response {
	var user model.User
	if service.PasswordConfirm != service.PasswordNew {
		return serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	user.ChangePassword(service.PasswordNew)
	model.DB.Save(&user)
	return serializer.BuildUserResponse(user)
}
