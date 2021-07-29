/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package article

import (
	"github.com/landers1037/blogo/models"
)

// DB_BLOG_LIKES 文章点赞表模型
type DB_BLOG_LIKES struct {
	models.Model
	Name string `gorm:"unique;not null" json:"name"`
	Like int `json:"like"`
}