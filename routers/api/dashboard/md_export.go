/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/config"
	"github.com/landers1037/blogo/models"
	"github.com/landers1037/blogo/models/admin"
	"github.com/landers1037/blogo/models/article"
	"github.com/landers1037/blogo/models/message"
	"github.com/landers1037/blogo/models/subscribe"
	"github.com/landers1037/blogo/utils"
	"net/http"
	"os/exec"
)

// ExportPosts 文章导出
// 分为全部导出和单个文章的导出
// 不传递参数时为全部导出
// 生成文件的逻辑由前端处理
func ExportPosts(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		var posts []article.DB_BLOG_POST
		e := models.BlogDB.Model(article.DB_BLOG_POST{}).Find(&posts).Error
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "export data fail",
				"data": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "export data success",
			"data": posts,
		})
		return
	}
	var post article.DB_BLOG_POST
	e := models.BlogDB.Model(article.DB_BLOG_POST{}).Where("name = ?", name).First(&post).Error
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "export data fail",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "export data success",
		"data": post,
	})
	return
}

// ExportDataBase 数据库导出
func ExportDataBase(c *gin.Context) {

}

// InitDB 数据库初始化
func InitDB(c *gin.Context) {
	models.BlogDB.CreateTable(admin.DB_BLOG_ADMIN{})
	models.BlogDB.CreateTable(article.DB_BLOG_POST{})
	models.BlogDB.CreateTable(article.DB_BLOG_VIEWS{})
	models.BlogDB.CreateTable(article.DB_BLOG_LIKES{})
	models.BlogDB.CreateTable(article.DB_BLOG_SHARE{})
	models.BlogDB.CreateTable(article.DB_BLOG_TAGS{})
	models.BlogDB.CreateTable(article.DB_BLOG_CATES{})
	models.BlogDB.CreateTable(article.DB_BLOG_ZHUANLAN{})
	models.BlogDB.CreateTable(article.DB_BLOG_COMMENTS{})
	models.BlogDB.CreateTable(message.DB_BLOG_MESSAGES{})
	models.BlogDB.CreateTable(subscribe.DB_BLOG_SUBSCRIBE{})
	c.JSON(http.StatusOK, gin.H{
		"msg": "db init",
		"data": "success",
	})
}

// BackUpDB 数据库备份
// 全局的备份路径在预定义里为/home/backup
// 防止io出错 使用shell完成
func BackUpDB(c *gin.Context) {
	date := utils.GetDate()
	cmd := fmt.Sprintf("cp %s /home/backup/blog%s.db",
		config.Cfg.DB, date)
	e := exec.Command("bash", "-c", cmd).Run()
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "backup db failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "backup db success",
		"data": "success",
	})
}