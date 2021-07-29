/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package message

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/models/dao/message_dao"
	"net/http"
)

type message struct {
	Mes string `json:"mes"`
}

// Savemes 保存留言
func Savemes(c *gin.Context)  {
	var m message
	err := c.BindJSON(&m)
	if err != nil{
		c.JSON(http.StatusForbidden,"error")
	}else {
		flag := message_dao.SaveMessage(m.Mes)
		if !flag{
			c.JSON(http.StatusOK,"error")
		}else {
			c.JSON(http.StatusOK,"saved")
		}

	}
}

// Getmes 获取留言
func Getmes(c *gin.Context) {
	data := message_dao.GetMessage()
	c.JSON(http.StatusOK,data)
}