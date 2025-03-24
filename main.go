package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	showOpening()
	mainLoop()
}

// 通用输入函数
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}

// 清屏（跨平台）
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// 等待回车
func waitForInput() {
	fmt.Print("\n按回车继续...")
	getInput()
}