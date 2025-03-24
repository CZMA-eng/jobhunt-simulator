package utils

import (
	"bufio"
	"fmt"
	"os"
)

// 通用输入函数
func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}

// 清屏（跨平台）
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// 等待回车
func WaitForInput() {
	fmt.Print("\n按回车继续...")
	GetInput()
}