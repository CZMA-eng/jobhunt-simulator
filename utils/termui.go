package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"

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

func PrintRainbowText(text string) {
	colors := []*color.Color{
		color.New(color.FgHiRed, color.BlinkSlow),    // 血红色闪烁
		color.New(color.FgHiGreen, color.Bold),       // 韭菜绿
		color.New(color.FgHiYellow, color.BgBlack),   // 福报黄
		color.New(color.FgHiBlue, color.Underline),   // 蓝图蓝
		color.New(color.FgHiMagenta, color.BlinkRapid),// 洗脑紫
	}

	rand.Seed(time.Now().UnixNano())
	for _, c := range text {
		// 随机字母大小写（模仿领导随意改需求）
		char := string(c)
		if rand.Intn(100) > 70 {
			char = string(c ^ 32) // 通过ASCII异或32实现大小写切换
		}

		// 随机颜色+动态打字效果
		colors[rand.Intn(len(colors))].Print(char)
		time.Sleep(time.Duration(50 + rand.Intn(50)) * time.Millisecond)
	}
	fmt.Println()
}

func PrintPoisonText(text string) {
	heartbeat := []rune{
		'▁', '▂', '▃', '▄', '▅', '▆', '▇', '█',
		'▇', '▆', '▅', '▄', '▃', '▂', '▁', ' ',
	}

	c := color.New(color.FgHiWhite, color.BgHiRed, color.BlinkRapid)
	for i := 0; ; i++ {
		// 生成随机偏移
		offset := rand.Intn(utf8.RuneCountInString(text)+3) - 1
		dirtyText := strings.Repeat(" ", offset) + text

		// 添加心电图动画
		c.Printf("%s %s", dirtyText, string(heartbeat[i%len(heartbeat)]))
		time.Sleep(200 * time.Millisecond)
		fmt.Print("\r\033[K") // 回退到行首并清空
	}
}

func PrintProgressBar(seconds int, label string) {
	// 伪代码生成器
	codeWords := []string{"敏捷", "赋能", "抓手", "闭环", "迭代", "痛点", "链路", "生态"}

	// 创建进度条协程
	done := make(chan bool)
	go func() {
		width := 50
		for i := 0; i <= width; i++ {
			// 生成随机伪代码
			var code strings.Builder
			for j := 0; j < 5; j++ {
				code.WriteString(codeWords[rand.Intn(len(codeWords))])
				code.WriteString(fmt.Sprintf("%d.0 ", rand.Intn(10)))
			}

			// 打印动态进度条
			fmt.Printf("\r%s [%s%s] %d%% %s",
				label,
				strings.Repeat("■", i),
				strings.Repeat(" ", width-i),
				i*2,
				code.String(),
			)
			time.Sleep(time.Duration(seconds)*time.Second/time.Duration(width))
		}
		done <- true
	}()

	// 添加闪烁光标
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				fmt.Print("\033[?25l") // 隐藏光标
				time.Sleep(500 * time.Millisecond)
				fmt.Print("\033[?25h") // 显示光标
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-done
	fmt.Println() // 换行结束
}
