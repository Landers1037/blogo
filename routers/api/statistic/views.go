/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package statistic

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/models/dao/statistics_dao"
	"net/http"
)

// GetViews 所有统计访问次数
func GetViews(c *gin.Context)  {
	uv := statistics_dao.StatisticViewQuery("all")
	c.JSON(http.StatusOK,uv)
}
