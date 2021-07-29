/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package logger

import (
	"fmt"
)

const logo = " ____  _     ____  _____\n/  __\\/ \\   /  _ \\/  __/\n| | //| |   | / \\|| |  _\n| |_\\\\| |_/\\| \\_/|| |_//\n\\____/\\____/\\____/\\____\\"

func PrintLogo() {
	for i:=0;i<=24;i++ {
		fmt.Printf("+")
	}
	fmt.Printf("\n%s\n", logo)
	for i:=0;i<=24;i++ {
		fmt.Printf("+")
	}
	fmt.Printf("\n\n")
}
