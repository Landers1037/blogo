/*
landers Apps
Author: landers
Github: github.com/landers1037

NAME:
   blogo service - 服务管理

USAGE:
   blogo service [command options] [arguments...]

CATEGORY:
   Service manager

OPTIONS:
   --conf value, -c value  指定配置文件路径 (default: "conf/app.ini")
   --help, -h              show help (default: false)
*/

package cmd

import (
	"fmt"
	"github.com/landers1037/blogo/utils/cmd"
	"github.com/landers1037/blogo/utils/settings"
	"github.com/urfave/cli/v2"
	"runtime"
)

// ServiceManager 服务管理动作
func ServiceManager() cli.ActionFunc {
	return func(c *cli.Context) error {
		s := c.String("s")
		// 基于配置文件 所以配置文件必须存在
		cf := c.String("c")

		avil1 := runtime.GOOS != "linux"
		avil2 := runtime.GOOS != "unix"
		avil3 := runtime.GOOS != "darwin"
		if avil1 && avil2 && avil3 {
			fmt.Println("你的操作系统不支持此特性")
			return nil
		}

		cfg, e := settings.InitCfg(cf)
		if e != nil {
			return e
		}
		switch s {
		case "start":
			cmd.StartAPP(cfg.AppName, cfg.APPPid, cfg.APPLog)
		case "stop":
			cmd.StopAPP(cfg.AppName, cfg.APPPid)
		case "restart":
			cmd.RestartAPP(cfg.AppName, cfg.APPPid, cfg.APPLog)
		case "status":
			cmd.StatusAPP(cfg.AppName, cfg.APPPid)
		default:
			return nil
		}
		return nil
	}
}

// AddServiceCmds 添加服务管理的命令
func AddServiceCmds() []*cli.Command {
	return []*cli.Command{
		{
			Name:                   "service",
			Aliases:                []string{"s"},
			Usage:                  "服务管理",
			Category:               "Service manager",
			Action: func(c *cli.Context) error {
				// 基于配置文件 所以配置文件必须存在
				cf := c.String("conf")
				s := c.Args().First()
				avil1 := runtime.GOOS != "linux"
				avil2 := runtime.GOOS != "unix"
				avil3 := runtime.GOOS != "darwin"
				if avil1 && avil2 && avil3 {
					fmt.Println("你的操作系统不支持此特性")
					return nil
				}
				cfg, e := settings.InitCfg(cf)
				if e != nil {
					return e
				}
				switch s {
				case "start":
					cmd.StartAPP(cfg.AppName, cfg.APPPid, cfg.APPLog)
				case "stop":
					cmd.StopAPP(cfg.AppName, cfg.APPPid)
				case "restart":
					cmd.RestartAPP(cfg.AppName, cfg.APPPid, cfg.APPLog)
				case "status":
					cmd.StatusAPP(cfg.AppName, cfg.APPPid)
				default:
					cmd.StatusAPP(cfg.AppName, cfg.APPPid)
				}
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "conf",
					Aliases: []string{"c"},
					Usage:   "指定配置文件路径",
					Value:   settings.APP_FILE,
				},
			},
		},
	}
}