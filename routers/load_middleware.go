/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/logger"
	"github.com/landers1037/blogo/middleware"
)

func LoadMiddleWare(r *gin.Engine) {
	logger.BlogLogger.InfoF("开始加载中间件")
	r.Use(middleware.BlogLogger())
	r.Use(middleware.Cors())
	r.Use(middleware.AllowIe())
	r.Use(middleware.TryFile())
	// 统计类的中间件分路由组编写
	r.Use(middleware.Uv())
	r.Use(middleware.PostView())
	// r.Use(middleware.SimpleAuth())
	r.Use(gin.Recovery())
	logger.BlogLogger.InfoF("中间件加载完毕")
}
