/*
Name: blog
File: update.go
Author: Landers
*/

package comment_dao

import (
	"github.com/landers1037/blogo/logger"
	"github.com/landers1037/blogo/models"
	"github.com/landers1037/blogo/models/article"
)

// CommentUpdate 评论的更新
func CommentUpdate(name string, id int, com string) error {
	e := models.BlogDB.Model(&article.DB_BLOG_COMMENTS{}).
		Where(map[string]interface{}{"name": name, "primary_id": id}).
		Update("comment", com).Error
	if e != nil {
		logger.BlogLogger.ErrorF("更新%s评论%s失败", name, id)
	}
	return e
}
