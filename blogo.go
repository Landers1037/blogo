/*
blogo is an application for creating your own blog with markdown
Author: landers
Github: github.com/landers1037

Usage:
NAME:
   blogo - powerful markdown-based blog

USAGE:
   blogo [global options] command [command options] [arguments...]

VERSION:
   v1.0.7

DESCRIPTION:
   一个基于markdown文档的动态博客部署工具

AUTHORS:
   Landers <liaorenj@gmail.com>
   wxk <xk_wang@qq.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command
   App configs:
     config, conf, c  应用配置
   Run a web service:
     web, w, serve, server  启动blog的web服务
   Service manager:
     service, s  服务管理
   Tools of blog:
     tool, t  博客配套工具

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   renj.io 2021.
 */
package main

import (
	"github.com/landers1037/blogo/cmd"
)

func main()  {
	cmd.RegisterCLI()
}