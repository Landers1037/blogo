/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: blog
*/

package sql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/landers1037/blogo/config"
)

var  (
	dbType, dbName, user, password, host, tablePrefix string
)

// Mysql mysql的连接池
func Mysql()  (db *gorm.DB, err error){
	//使用mysql
	dbType = config.Cfg.MySQLType
	dbName = config.Cfg.MySQLName
	user = config.Cfg.MySQLUserName
	password = config.Cfg.MySQLPassWd
	host = config.Cfg.MySQLHost
	tablePrefix = config.Cfg.MySQLTablePrefix

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	return
}



