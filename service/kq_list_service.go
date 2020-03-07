package service

import (
	"swuxfdc/model"
	"swuxfdc/serializer"
)

// ListKqService 考勤表列表服务
type ListKqService struct {
	Limit   int    `form:"limit"`
	Start   int    `form:"start"`
	Grade   string `form:"q_grade"`
	Classes string `form:"q_classes"`
	Week    int    `form:"q_week"`
}

// List 考勤表列表
func (service *ListKqService) List(user model.User) serializer.Response {
	kqs := []model.Kq{}
	total := 0
	var query model.Kq
	query.Week = service.Week
	query.Grade = service.Grade
	query.Classes = service.Classes

	if user.Dept == "" {
		query.Author = user.Name
	}

	if service.Limit == 0 {
		service.Limit = 100
	}

	if err := model.DB.Model(model.Kq{}).Where(query).Count(&total).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Order("id desc", true).Where(query).Find(&kqs).Error; err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildKqs(kqs), uint(total))

}
