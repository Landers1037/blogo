/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package config

import (
	"github.com/landers1037/blogo/utils/settings"
	"log"
	"os"
)

var Cfg settings.Cfg

func InitCfg(p string) {
	var e error
	Cfg, e = settings.InitCfg(p)
	if e != nil && os.IsNotExist(e) {
		err := settings.SaveConfig(p)
		if err != nil {
			log.Fatalf("程序配置文件初始化失败 %s", e.Error())
		}
		// 通过结构体重载配置
		Cfg, e = settings.InitCfg(p)
		if e != nil {
			log.Fatalf("程序配置文件读取失败 %s", e.Error())
		}
	}
}
