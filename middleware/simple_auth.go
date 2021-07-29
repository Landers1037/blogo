/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package middleware
//基于referer的简单验证

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/config"
	"strings"
)

// SimpleAuth 简单校验只允许特定refer和host通过
func SimpleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var referer= c.Request.Referer()
		var host = c.Request.Host
		allowRefer := config.Cfg.AppRefer
		allowHost := config.Cfg.AppHost

		ifRefer := check(strings.Fields(allowRefer), referer)
		ifHost := check(strings.Fields(allowHost), host)

		if ifRefer && ifHost{
			// 处理请求
			c.Next()
		}else {
			c.AbortWithStatus(403)
			return
		}

	}
}

func check(strList []string, sub string) bool {
	for _, v := range strList {
		if strings.Contains(sub, v) {
			return true
		}
	}
	return false
}