/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/landers1037/blogo/models/dao/dberr"
	"github.com/landers1037/blogo/models/dao/post_dao"
	"github.com/landers1037/blogo/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// UploadFile markdown文件的上传解析
// 因为不会保存为临时文件 所以直接读取上传的字节流
func UploadFile(c *gin.Context) {
	file, _, e := c.Request.FormFile("uploadmd")
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file upload failed",
			"data": "fail",
		})
		return
	}
	fileData, e := ioutil.ReadAll(file)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file parse failed",
			"data": "fail",
		})
		return
	}
	if ok := utils.CheckArticleOK(fileData);!ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file check failed",
			"data": "fail",
		})
		return
	}else {
		// 开始解析md字节流
		mdData, e := utils.GenMdData(fileData)
		if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "file parse failed",
				"data": "fail",
			})
			return
		}
		e = post_dao.PostAdd(mdData)
		// 可能是更新文章所以判断错误是否是因为重复
		if e != nil && e == dberr.ERR_ARTICLE_EXIST {
			e = post_dao.PostUpdateAll(mdData)
			if e != nil {
				c.JSON(http.StatusOK, gin.H{
					"msg": "file upload failed",
					"data": "fail",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "file upload success",
				"data": "success",
			})
			return
		}else if e != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "file upload failed",
				"data": "fail",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "file upload success",
			"data": "success",
		})
		return
	}
}

// CheckFile 校验md合法性
func CheckFile(c *gin.Context) {
	file, _, e := c.Request.FormFile("uploadmd")
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file upload failed",
			"data": "fail",
		})
	}
	fileData, e := ioutil.ReadAll(file)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file parse failed",
			"data": "fail",
		})
	}
	if ok := utils.CheckArticleOK(fileData);ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file check success",
			"data": "success",
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file check failed",
			"data": "fail",
		})
	}
}

// UploadFileCallBack 上传回调 返回解析的内容
func UploadFileCallBack(c *gin.Context) {
	file, _, e := c.Request.FormFile("uploadmd")
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file upload failed",
			"data": "fail",
		})
		return
	}
	fileData, e := ioutil.ReadAll(file)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file parse failed",
			"data": "fail",
		})
		return
	}
	meta, e := utils.ParseYamlFront(fileData)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "file parse failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "file parse success",
		"data": meta,
	})
}

type newPostData struct {
	Name string
	Title string
	Tags string
	Date string
}

// NewPost 在管理面直接新增一片空白的文章 只有name是需要的title默认等于name
func NewPost(c *gin.Context)  {
	var np newPostData
	e := c.BindJSON(&np)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "data parse failed",
			"data": "fail",
		})
		return
	}
	var mdData utils.MdData
	mdData.Meta.Name = np.Name
	mdData.Meta.Date = strings.Split(np.Date, " ")[0]
	mdData.Meta.DatePlus = np.Date
	mdData.Meta.Title = np.Title
	mdData.Meta.Tags = strings.Fields(np.Tags)
	e = post_dao.PostAdd(mdData)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "post add failed",
			"data": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "post add success",
		"data": "success",
	})
	return
}