package model

import (
	"github.com/jinzhu/gorm"
)

// Kq 用户模型
type Kq struct {
	gorm.Model
	Grade   string //年级
	Classes string //班级
	Week    int    //第几周
	Day     int    //星期几
	KqURL   string //图片链接
}
