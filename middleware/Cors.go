/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/config"
	"net/http"
)

// Cors 开启跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.Cfg.CORSFlag {
			method := c.Request.Method
			// 使用cookie时 不能使用*
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, admin_token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

			//放行所有OPTIONS方法
			if method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
			// 处理请求
			c.Next()
		}
	}
}

