/*
Name: blog
File: create.go
Author: Landers
*/

package share_dao

import (
	"github.com/landers1037/blogo/models"
	"github.com/landers1037/blogo/models/article"
	"sync"
)

var shareLock sync.Mutex

// ShareAdd 分享量增加
func ShareAdd(name string) error {
	var s article.DB_BLOG_SHARE
	e := models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Where("name = ?", name).First(&s).Error
	if e != nil {
		s.Share = 1
		s.Name = name
		return models.BlogDB.Model(&article.DB_BLOG_SHARE{}).Create(&s).Error
	}else {
		shareLock.Lock()
		defer shareLock.Unlock()
		s.Share = s.Share + 1
		return models.BlogDB.Save(&s).Error
	}
}
