/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package article

import "github.com/landers1037/blogo/models"

// DB_BLOG_TAGS 文章标签表模型
type DB_BLOG_TAGS struct {
	models.Model
	Tag string `json:"tag"`
	Name string `json:"name"`
}
