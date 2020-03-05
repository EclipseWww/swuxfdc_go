package serializer

import "swuxfdc/model"

// Kq 考勤表序列化器
type Kq struct {
	CreatedAt int64  `json:"created_at"`
	Grade     string `json:"grade"`   //年级
	Classes   string `json:"classes"` //班级
	Week      int    `json:"week"`    //第几周
	Day       int    `json:"day"`     //星期几
	KqURL     string `json:"kq_URL"`  //图片链接
}

// BuildKq 序列化考勤表
func BuildKq(item model.Kq) Kq {
	return Kq{
		Week:      item.Week,
		Day:       item.Day,
		Grade:     item.Grade,
		Classes:   item.Classes,
		KqURL:     item.KqURL,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
