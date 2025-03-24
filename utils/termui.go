package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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

// Confirm 显示确认对话框，返回用户是否选择"是"
func Confirm(prompt string) bool {
    fmt.Printf("%s (y/N) > ", prompt)
    input := strings.ToLower(strings.TrimSpace(GetInput()))
    return input == "y" || input == "yes"
}

// ShowLoading 显示加载进度条
// seconds - 总时长(秒)
// message - 加载时显示的消息
func ShowLoading(seconds int, message string) {
    const width = 50 // 进度条宽度
    interval := time.Duration(float64(seconds)*1000/float64(width)) * time.Millisecond

    fmt.Printf("%s [%s]", message, strings.Repeat(" ", width))
    fmt.Printf("\r%s [", message) // 回到行首

    for i := 0; i < width; i++ {
        fmt.Print("█")
        time.Sleep(interval)
        
        // 显示百分比
        percent := (i + 1) * 100 / width
        fmt.Printf("] %d%%", percent)
        if i < width-1 {
            fmt.Printf("\r%s [", message) // 回到进度条开始位置
        }
    }
    fmt.Println() // 最终换行
}