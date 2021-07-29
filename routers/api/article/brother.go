/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package article

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/models/dao/post_dao"
	"net/http"
)

// Getbrother 处理上下篇
func Getbrother(c *gin.Context)  {
	name := c.Query("name")
	p,n := post_dao.Getbrother(name)
	var data = []string{p,n}
	c.JSON(http.StatusOK,data)
}
