package randomEvents

import (
	"jobhunt/players"
	"jobhunt/utils"
	"math/rand"

	"github.com/fatih/color"
)

// 随机事件系统
func RandomEvent(p *players.Player) {
	events := []func(*players.Player){
		// 正能量事件（极少）
		func(p *players.Player) {
			if rand.Intn(100) > 95 { // 5%概率
				color.HiGreen("\n[系统] 收到意外赞美！")
				color.White("某HR在后台备注：")
				color.Green("'这人简历虽然菜但照片挺好看'")
				p.Hope += 20
			}
		},
		// 负能量事件
		func(p *players.Player) {
			color.Red("\n[紧急] 班级群通知：")
			color.White("你的室友 %s 拿到了%soffer",
				color.HiGreenString("张三"),
				color.HiYellowString("50万/年"))
			p.Hope -= 25
		},
		func(p *players.Player) {
			color.HiRed("\n[推送] 知乎热榜：")
			color.White("《为什么说35岁程序员不如狗》")
			color.HiBlack("阅读 10w+ · 收藏 5w+")
			p.Sanity -= 15
		},
	}
	
	// 80%概率触发负事件
	if rand.Intn(100) > 20 {
		events[1+rand.Intn(len(events)-1)](p)
	} else {
		events[0](p)
	}
	utils.WaitForInput()
}