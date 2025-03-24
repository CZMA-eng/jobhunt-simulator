package boss

import (
	"fmt"
	"github.com/fatih/color"
	"jobhunt/players"
	"jobhunt/utils"
	"time"
)

// 进入BOSS直聘模式
func EnterBossMode(p *players.Player) {
	utils.ClearScreen()
	
	// 初始化压力系统
	pressure := startPressureSystem(p)
	defer pressure.stop()
	
	// 显示模式标题
	showBossHeader()
	
	// 加载连接动画
	utils.ShowLoadingWithLabel(2, "正在连接暗网招聘系统...", 20)
	
	// 触发随机事件
	triggerRandomEvent(p)
	
	// 检查成就解锁
	checkAchievements(p)
	
	utils.WaitForInput()
}

func showBossHeader() {
	// 创建一个红色打印器
	redPrinter := color.New(color.FgHiRed)
	redPrinter.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	redPrinter.Println("▓▓                        ▓▓")
	redPrinter.Println("▓▓   BOSS直聘 - 地狱模式   ▓▓")
	redPrinter.Println("▓▓                        ▓▓")
	redPrinter.Println("▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println()
}

// 压力系统相关结构
type pressureSystem struct {
	ticker   *time.Ticker
	done     chan struct{}
	pressure int
}

func startPressureSystem(p *players.Player) *pressureSystem {
	ps := &pressureSystem{
		ticker: time.NewTicker(8 * time.Second),
		done:   make(chan struct{}),
	}
	
	go func() {
		// 创建颜色打印器（黑色高亮）
		pressurePrinter := color.New(color.FgHiBlack)
		
		for {
			select {
			case <-ps.ticker.C:
				ps.pressure++
				pressurePrinter.Printf("[压力辐射] 你的焦虑值 +%d\n", ps.pressure)
				p.Sanity -= ps.pressure * 3
				p.Hope -= ps.pressure
				
			case <-ps.done:
				return
			}
		}
	}()
	return ps

}

func (ps *pressureSystem) stop() {
	ps.ticker.Stop()
	close(ps.done)
}