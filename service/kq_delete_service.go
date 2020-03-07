package service

import (
	"swuxfdc/model"
	"swuxfdc/serializer"
)

// DeleteKqService 删除考勤表的服务
type DeleteKqService struct {
}

// Delete 删除考勤表
func (service *DeleteKqService) Delete(id string, name string) serializer.Response {
	var kq model.Kq
	err := model.DB.First(&kq, id).Error

	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "考勤表不存在",
			Error: err.Error(),
		}
	}
	if name != kq.Author {
		return serializer.Response{
			Code:  404,
			Msg:   "无此考勤表删除权限",
			Error: "",
		}
	}
	err = model.DB.Delete(&kq).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "考勤表删除失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{}
}
