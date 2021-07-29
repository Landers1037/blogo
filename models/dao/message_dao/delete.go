/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"github.com/landers1037/blogo/models"
	"github.com/landers1037/blogo/models/message"
)

// DelMessage 删除留言
func DelMessage(id int) error {
	// 尽力删除模型
	e := models.BlogDB.Delete(&message.DB_BLOG_MESSAGES{}, "primary_id = ?", id).Error
	return e
}