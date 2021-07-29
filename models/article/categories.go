/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"github.com/landers1037/blogo/models"
)

// DB_BLOG_CATES 文章分类表模型
type DB_BLOG_CATES struct {
	models.Model
	Cate string `json:"cate"`
	Name string `json:"name"`
}