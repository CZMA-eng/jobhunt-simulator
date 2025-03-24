package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
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

func FakeLoading(seconds int) {
	fmt.Print("[")
	for i := 0; i < 20; i++ {
		fmt.Print(" ")
	}
	fmt.Print("]\r[")
	
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(seconds)*50 * time.Millisecond)
		fmt.Print("█")
	}
}

// ShowLoadingWithLabel 带标签的加载条
func ShowLoadingWithLabel(seconds int, label string, width int) {
	interval := time.Duration(float64(seconds)*1000/float64(width)) * time.Millisecond
	fmt.Printf("%s [%s]", label, strings.Repeat(" ", width))
	fmt.Printf("\r%s [", label)
	
	for i := 0; i < width; i++ {
		fmt.Print("█")
		time.Sleep(interval)
	}
	fmt.Println("]")
}

// PrintBulletList 打印项目符号列表
func PrintBulletList(items []string, bulletColor color.Attribute) {
	c := color.New(bulletColor)
	for _, item := range items {
		c.Print("• ")
		fmt.Println(item)
	}
}

// PlaySoundEffect 音效播放（模拟实现）
func PlaySoundEffect(name string) {
	switch name {
	case "achievement_unlock":
		fmt.Print("\a") // 系统提示音
	default:
		fmt.Print("\a")
	}
}

// TypewriterEffect 打字机效果
func TypewriterEffect(text string, delayMs int) {
	for _, c := range text {
		fmt.Print(string(c))
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}

func ShowDynamicInputIndicator(seconds int) {
	frames := []string{"", ".", "..", "..."}
	endTime := time.Now().Add(time.Duration(seconds) * time.Second)
	
	for time.Now().Before(endTime) {
		for _, frame := range frames {
			fmt.Printf("\r对方正在输入%s   ", frame)
			time.Sleep(300 * time.Millisecond)
		}
	}
	fmt.Print("\r")
}