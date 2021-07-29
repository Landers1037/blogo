/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/landers1037/blogo/models/article"
	"strings"
)

// ImageMigrate 专用于图片资源的地址迁移工作
func ImageMigrate(old, new string) error {
	var posts []article.DB_BLOG_POST
	db, e := gorm.Open("sqlite3", "blog.db")
	if e != nil {
		return e
	}
	db.SingularTable(true)
	// 默认会自动对全部的文章遍历
	e = db.Model(&article.DB_BLOG_POST{}).Find(&posts).Error
	if e != nil {
		return e
	}
	// 使用strings自动替换全部符合条件的字符串
	for _, p := range posts {
		fmt.Printf("开始迁移%s\n", p.Name)
		var oldAbs, oldContent string
		oldAbs = strings.ReplaceAll(p.Abstract, old, new)
		oldContent = strings.ReplaceAll(p.Content, old, new)
		p.Abstract = oldAbs
		p.Content = oldContent
		e := db.Save(&p).Error
		if e != nil {
			fmt.Printf("迁移%s失败: %s\n", p.Name, e.Error())
		}
	}
	return nil
}
