/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: blog
*/

package models
//数据库的初始化
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/landers1037/blogo/logger"
	"github.com/landers1037/blogo/models/sql"
	_ "github.com/mattn/go-sqlite3"
)

var BlogDB *gorm.DB

// InitDB 数据库连接初始化
func InitDB() {
	var err error
	// 使用sqlite
	BlogDB, err = sql.Sqlite()
	if err != nil {
		logger.BlogLogger.ErrorF("数据库连接失败 %s", err.Error())
		logger.BlogLogger.PanicF("数据库错误 异常退出")
	}

	BlogDB.SingularTable(true)
	BlogDB.DB().SetMaxIdleConns(10)
	BlogDB.DB().SetMaxOpenConns(100)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	logger.BlogLogger.InfoF("数据库连接已关闭")
	if BlogDB != nil {
		defer BlogDB.Close()
	}
}
