/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/utils/err"
	"net/http"
)

// AdminLogout 告知前端清空全部cookie
// 清空cookie的操作交给前端 后台只需响应
func AdminLogout(c *gin.Context) {
	code := err.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"code": code,
		"msg": err.GetMsg(code),
		"data": "success",
	})
}
