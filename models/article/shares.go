/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"github.com/landers1037/blogo/models"
)

// DB_BLOG_SHARE 文章分享表模型
type DB_BLOG_SHARE struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	Share int `json:"share"`
}