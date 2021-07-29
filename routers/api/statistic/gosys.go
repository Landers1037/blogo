/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package statistic

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/utils/cmd"
	"net/http"
)

func GetRoutines(c * gin.Context)  {
	n := cmd.GetGID()

	c.JSON(http.StatusOK,n)
}

func GetMem(c *gin.Context)  {
	o := cmd.Sh("mem")

	c.JSON(http.StatusOK,o)
}


