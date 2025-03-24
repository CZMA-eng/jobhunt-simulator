package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// QuitWithStyle 带嘲讽的退出函数
func QuitWithStyle() {
    ClearScreen()

    // 随机选择一种退出风格
    switch rand.Intn(5) {
    case 0:
        corporateQuit()
    case 1:
        hackerQuit()
    case 2:
        existentialQuit()
    case 3:
        passiveAggressiveQuit()
    default:
        brutalHonestyQuit()
    }

    time.Sleep(2 * time.Second)
    os.Exit(0)
}

// ========== 私有退出动画 ==========

func corporateQuit() {
    color.HiBlue(`
    ╔════════════════════════╗
    ║                        ║
    ║  感谢您使用求职模拟器  ║
    ║                        ║
    ║ 退出原因代码:           ║
    ║   • 0xDEADBEEF         ║
    ║   • 心理承受力不足      ║
    ║                        ║
    ╚════════════════════════╝
    `)
    fmt.Println(color.HiBlackString("(你的表现已计入HR黑名单数据库)"))
}

func hackerQuit() {
    fmt.Println(color.HiGreenString("> ./jobhunt_simulator --quit --reason=weak_mental"))
    fmt.Println(color.HiBlackString("Segmentation fault (core dumped)"))
    fmt.Println()
    
    frames := []string{
        "|", "/", "-", "\\",
        "|", "/", "-", "\\",
    }
    
    for i := 0; i < 3; i++ {
        for _, frame := range frames {
            fmt.Printf("\r[%s] 正在删除你的求职信心...", frame)
            time.Sleep(100 * time.Millisecond)
        }
    }
    fmt.Println("\n\n" + color.HiRedString("删除成功！"))
}

func existentialQuit() {
    color.HiMagenta(`
    ╔════════════════════════════════╗
    ║                                ║
    ║  你确定要回到"现实"吗？         ║
    ║                                ║
    ║   • 这里至少不用真的被HR拒绝    ║
    ║   • 虚拟痛苦 vs 真实痛苦       ║
    ║   • 404: 人生意义未找到        ║
    ║                                ║
    ╚════════════════════════════════╝
    `)
    fmt.Println(color.HiYellowString("\n(系统提示：你的存在本身就是一个bug)"))
}

func passiveAggressiveQuit() {
    color.HiYellow(`
    ╭───────────────────────────╮
    │                           │
    │  我们很遗憾看到你离开...  │
    │                           │
    │  其他求职者平均承受时间：  │
    │   • 比你长3倍             │
    │   • 收到拒信多5倍         │
    │   • 但坚持下来了          │
    │                           │
    ╰───────────────────────────╯
    `)
    fmt.Println(color.HiBlackString("\n(已自动将你的简历标记为'玻璃心'类别)"))
}

func brutalHonestyQuit() {
    color.HiRed(`
    ▄︻デ══━ 求职模拟终止 ━══デ︻▄
    
    坦白说：
    • 你连虚拟求职都撑不住
    • 现实中的HR会更残忍
    • 你的专业正在被AI取代
    • 你的简历会被自动过滤
    
    建议：
    ` + color.HiGreenString("转行送外卖") + `
    `)
    fmt.Println(color.HiBlackString("\n(本消息由AI求职顾问生成)"))
}

// QuitConfirm 带特效的退出确认
func QuitConfirm(prompt string) bool {
	// 第一阶段：红色闪烁警告
	for i := 0; i < 3; i++ {
		ClearLine()
		color.HiBlack("> ")
		color.HiRed("%s", prompt)
		time.Sleep(200 * time.Millisecond)
		ClearLine()
		color.HiBlack("> ")
		color.HiYellow("%s", prompt)
		time.Sleep(200 * time.Millisecond)
	}

	// 第二阶段：动态输入提示
	fmt.Print(color.HiRedString("⚠️ "))
	color.HiBlack("["+time.Now().Format("15:04:05")+"] ")
	fmt.Print(color.HiRedString(prompt))
	
    fmt.Println()
    fmt.Printf(color.HiRedString("\r%s [y/N]", prompt))

	// 获取输入
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))


	return input == "y" || input == "yes"
}


// ClearLine 清除当前行
func ClearLine() {
	fmt.Print("\033[2K\r")
}
