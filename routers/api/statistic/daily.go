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

func GetDaily(c *gin.Context)  {
	count := statistics_dao.StatisticDailyQuery()
	c.JSON(http.StatusOK,count)
}
