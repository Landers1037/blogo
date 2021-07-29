/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package message_dao

import (
	"github.com/landers1037/blogo/logger"
	"github.com/landers1037/blogo/models"
	"github.com/landers1037/blogo/models/message"
	"github.com/landers1037/blogo/utils"
	"time"
)

// SaveMessage 添加留言
func SaveMessage(mes string) bool {
	m := message.DB_BLOG_MESSAGES{Message:mes, Date: utils.GetDatePlus()}
	e := models.BlogDB.Create(&m).Error
	logger.BlogLogger.ErrorF("%s", e)
	time.Sleep(1)
	if e != nil {
		return false
	}
	return true
}