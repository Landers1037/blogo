/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"github.com/landers1037/blogo/models"
)

// DB_BLOG_COMMENTS 文章评论表模型
type DB_BLOG_COMMENTS struct {
	models.Model
	Name string `json:"name"`
	User string `json:"user"`
	Date string `json:"date"`
	Comment string `json:"comment"`
}