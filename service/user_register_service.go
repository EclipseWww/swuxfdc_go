package service

import (
	"swuxfdc/model"
	"swuxfdc/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Name            string `form:"name" json:"name" `
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
	Grade           string `form:"grade" json:"grade" binding:"required,min=2,max=30"`
	Classes         string `form:"classes" json:"classes" binding:"required,min=2,max=30"`
	Dept            string `form:"dept" json:"dept" `
	Level           int    `form:"evel" json:"level" `
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("name = ?", service.Name).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "该姓名账户已存在",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Name:     service.Name,
		UserName: service.UserName,
		Status:   model.Active,
		Grade:    service.Grade,
		Classes:  service.Classes,
		Dept:     service.Dept,
		Level:    service.Level,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
