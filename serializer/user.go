package serializer

import "swuxfdc/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Grade     string `json:"grade"`   //年级
	Classes   string `json:"classes"` //班级
	Dept      string `json:"dept"`    //部门
	Level     int    `json:"level"`   //职位

}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Name:      user.Name,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Grade:     user.Grade,
		Classes:   user.Classes,
		Dept:      user.Dept,
		Level:     user.Level,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
