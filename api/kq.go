package api

import (
	"swuxfdc/service"

	"github.com/gin-gonic/gin"
)

// CreateKq 考勤表上传
func CreateKq(c *gin.Context) {
	user := CurrentUser(c)
	service := service.UploadKqService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Upload(user.Name)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteKq 删除考勤表的接口
func DeleteKq(c *gin.Context) {
	user := CurrentUser(c)
	service := service.DeleteKqService{}
	res := service.Delete(c.Param("id"), user.Name)
	c.JSON(200, res)
}

// ListKq 考勤表列表接口
func ListKq(c *gin.Context) {
	user := CurrentUser(c)
	service := service.ListKqService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(*user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//这段代码用来获取上下文参数 从interface转换为*model.User 直接调用CurrentUser获取用户
// var test1 *model.User
// test, err := c.Get("user")
// test1 = test.(*model.User)
// fmt.Print(err, test1.Classes)
// c.JSON(200, test)

// 类型转换userIDString := fmt.Sprint(userID)
