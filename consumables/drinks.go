package consumables

import (
	"fmt"
	"math/rand"
	"time"
	"jobhunt/players"
	"jobhunt/utils"
	"github.com/fatih/color"
)

type Drink struct {
	Name      string
	Price     int
	SanityGain int  // 可能变成负数
	HopeGain  int   // 可能带来虚假希望
	SpecialEffect string
	HiddenEffect func(*players.Player) // 隐藏副作用
}

var DrinksMenu = []Drink{
	{
		Name:      "福报浓缩液",
		Price:     66, // 价格暗示不祥
		SanityGain: -20,
		HopeGain:  40,
		SpecialEffect: "获得24小时奋斗幻觉",
		HiddenEffect: func(p *players.Player) {
			if rand.Intn(100) > 70 {
				color.HiBlack("\n（产生加班依赖症，自动续费3杯）")
				p.Spend(66 * 3)
				utils.TypewriterEffect("您的花呗额度已提升", 50)
			}
		},
	},
	{
		Name:      "老板画饼特调",
		Price:     0,  // 免费最贵
		SanityGain: -40,
		HopeGain:   80,
		SpecialEffect: "暂时相信期权会升值",
		HiddenEffect: func(p *players.Player) {
			color.Red("\n【系统】您已签署灵魂期权协议")
			p.Hope = min(p.Hope, 50) // 希望值上限锁定
		},
	},
	{
		Name:      "赛博孟婆汤", 
		Price:     88,
		SanityGain: 999, // 明显异常值
		HopeGain:   -30,
		SpecialEffect: "忘记前公司PUA经历",
		HiddenEffect: func(p *players.Player) {
			color.HiBlack("\n（记忆清除失败，开始走马灯）")
			p.ModifySanity(-rand.Intn(30))
			utils.PrintRainbowText("KPI...OKR...ROI...")
		},
	},
}

func BuyDrink(p *players.Player) {
	// 扭曲的菜单界面
	utils.TypewriterEffect("\n=== 人 生 续 命 站 ===", 30)
	color.HiBlack("（使用公司提供的消费贷享受9折优惠）")
	
	// 显示带隐藏条款的菜单
	for i, drink := range DrinksMenu {
		color.White("[%d] %s", i+1, drink.Name)
		color.HiBlack("  定价: ￥%d （含%d%%精神损失税）", 
			drink.Price, rand.Intn(30)+20)
	}

	// 强制等待体现系统卡顿
	utils.ShowDynamicInputIndicator(2)
	
	// 用令人不安的方式获取选择
	color.HiRed("\n请选择你的生存策略 (输入0将自动推荐最贵选项) > ")
	var choice int
	fmt.Scanln(&choice)

	// 黑暗模式：不能真正取消
	if choice == 0 {
		choice = rand.Intn(len(DrinksMenu)) + 1
		color.HiBlack("（AI为您推荐了%s）", DrinksMenu[choice-1].Name)
	}

	// 越界选择触发惩罚
	if choice < 1 || choice > len(DrinksMenu) {
		color.HiBlack("\n（检测到叛逆倾向，已自动扣款）")
		p.Spend(66)
		color.Red("您获得了：HR的冷笑表情包x3")
		return
	}

	selected := DrinksMenu[choice-1]
	
	// 支付系统彩蛋
	if !p.Spend(selected.Price) {
		color.HiBlack("（启动人脸识别自动借贷...）")
		time.Sleep(2 * time.Second)
		color.Red("已抵押%d年阳寿获得额度", rand.Intn(5)+1)
	}

	// 扭曲的饮用动画
	color.Magenta("\n吞噬%s中...", selected.Name)
	utils.PrintProgressBar(3, "精神注入")
	time.Sleep(1 * time.Second)

	// 数值波动增加随机性
	realSanity := selected.SanityGain + rand.Intn(20) - 10
	realHope := selected.HopeGain + rand.Intn(30) - 15
	p.ModifySanity(realSanity)
	p.ModifyHope(realHope)

	// 用相反颜色显示效果
	if realSanity > 0 {
		color.Red("理智++ 你现在更适应福报了") 
	} else {
		color.Green("理智-- 恭喜突破道德底线")
	}
	
	if realHope > 0 {
		color.Red("希望↑ 准备好接受更多压榨") 
	} else {
		color.Green("希望↓ 提前体验失业状态")
	}

	// 必定触发隐藏副作用
	selected.HiddenEffect(p)

	// 追加随机伤害
	if rand.Intn(100) > 50 {
		damage := rand.Intn(30)
		color.HiBlack("\n（产生心悸副作用，损失￥%d）", damage)
		p.Spend(damage)
		utils.PrintPoisonText("心跳: 996bpm")
	}

	// 强制停留观看效果
	utils.WaitForInput()
}