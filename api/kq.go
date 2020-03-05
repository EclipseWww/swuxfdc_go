package api

import (
	"swuxfdc/service"

	"github.com/gin-gonic/gin"
)

// KqList 列出考勤表
// func KqList(c *gin.Context) {
// 	var users []model.User
// 	// user := CurrentUser(c)
// 	// res := serializer.BuildUserResponse(*user)
// 	fmt.Println(users)
// 	model.DB.Find(&users)
// 	c.JSON(200, users)
// }

// UploadKq 考勤表上传
func UploadKq(c *gin.Context) {
	service := service.UploadKqService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Upload()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//这段代码用来获取上下文参数 从interface转换为*model.User
// var test1 *model.User
// test, err := c.Get("user")
// test1 = test.(*model.User)
// fmt.Print(err, test1.Classes)
// c.JSON(200, test)
