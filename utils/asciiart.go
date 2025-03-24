package utils

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func ShowOpening() {
	ClearScreen()
	
	// 破碎ASCII艺术字
	fig := figure.NewFigure("JobHunt 2024", "doom", true)
	fig.Print()
	
	fmt.Println()
	color.Red("■ 当前就业市场指标：")
	color.Yellow("  - 招聘需求同比下降 200%")
	color.Yellow("  - 平均每个岗位 114514 位竞争者")
	color.Yellow("  - 你的专业已被AI淘汰")
	
	fmt.Println()
	color.HiBlack("开始你的福报之旅...")
	WaitForInput()
}

func GameOver() {
	ClearScreen()
	
	endings := []string{
		`你成为了：
██████ 三和大神 ██████
- 日结工资系统已解锁
- 永久失去社保资格`,
		
		`结局：赛博精神病
✓ 成功识别所有JD黑话
✓ 患上HRPTSD
✓ 看见‘赋能’就想吐`,
	}
	
	color.HiMagenta(endings[rand.Intn(len(endings))])
	fmt.Println()
	color.HiBlack("（重新运行程序可再来一轮）")
	os.Exit(0)
}