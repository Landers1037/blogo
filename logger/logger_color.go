/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package logger

// 基于ansi的终端颜色
const (
	ERROR = 31
	INFO = 36
	PANIC = 35
)

const FG = ""
const BG = ""
const Fmt = "[\x1b[0;%dm%s\x1b[0m] %s\n"