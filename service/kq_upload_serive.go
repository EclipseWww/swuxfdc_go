package service

import (
	"swuxfdc/model"
	"swuxfdc/serializer"
)

// UploadKqService 考勤表上传的服务
type UploadKqService struct {
	Grade   string `form:"grade" json:"grade" binding:"required"`     //年级
	Classes string `form:"classes" json:"classes" binding:"required"` //班级
	Week    int    `form:"week" json:"week" binding:"required"`       //第几周
	Day     int    `form:"day" json:"day" binding:"required"`         //星期几
	KqURL   string `form:"kq_URL" json:"kq_URL" binding:"required"`   //图片链接
}

// Upload 考勤表上传
func (service *UploadKqService) Upload() serializer.Response {
	kq := model.Kq{
		Grade:   service.Grade,
		Classes: service.Classes,
		KqURL:   service.KqURL,
		Week:    service.Week,
		Day:     service.Day,
	}

	err := model.DB.Create(&kq).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "考勤表保存失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildKq(kq),
	}
}
