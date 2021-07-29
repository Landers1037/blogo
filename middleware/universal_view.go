/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/config"
	"github.com/landers1037/blogo/models/dao/statistics_dao"
	"github.com/landers1037/blogo/timer"
	"time"
)

// Uv 防止内存碎片化 使用按时写入机制
// 每600s刷新一次
func Uv() gin.HandlerFunc {
	return func(c *gin.Context) {
		//设计一个缓存的栈
		// 处理请求
		if config.Cfg.UvFlag{
			timer.UpdateUv()
			if time.Now().Minute() % 5 == 0 {
				statistics_dao.StatisticsViewUpdate("all", timer.GetUv())
				timer.ClearUv()
			}
		}
		c.Next()
	}
}

