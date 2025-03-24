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
	SanityGain int
	HopeGain  int
	SpecialEffect string
}

var DrinksMenu = []Drink{
	{
		Name:      "珍珠奶茶",
		Price:     25,
		SanityGain: 10,
		HopeGain:  5,
		SpecialEffect: "暂时忘记求职烦恼",
	},
	{
		Name:      "香菜柠檬特饮",
		Price:     15,
		SanityGain: -5, // 可能降低理智
		HopeGain:   20,  // 但给虚假希望
		SpecialEffect: "猎奇体验带来短暂快乐",
	},
	{
		Name:      "佛系养生茶",
		Price:     40,
		SanityGain: 20,
		HopeGain:   0,
		SpecialEffect: "看淡一切",
	},
}

func BuyDrink(p *players.Player) {
	// 显示菜单
	color.Cyan("\n=== 奶茶续命菜单 ===")
	for i, drink := range DrinksMenu {
		fmt.Printf("[%d] %s ￥%d\n", i+1, 
			color.HiWhiteString(drink.Name), 
			drink.Price)
		fmt.Printf("   效果: 理智+%d 希望+%d | %s\n",
			drink.SanityGain, drink.HopeGain,
			color.HiBlackString(drink.SpecialEffect))
	}

	// 选择
	fmt.Print("\n选择饮品 (0取消) > ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(DrinksMenu) {
		color.Red("无效选择！")
		return
	}

	selected := DrinksMenu[choice-1]
	
	// 支付检查
	if !p.Spend(selected.Price) {
		return
	}

	// 应用效果
	color.Magenta("\n你喝下了%s...", selected.Name)
	time.Sleep(1 * time.Second)

	p.Sanity = min(100, p.Sanity+selected.SanityGain)
	p.Hope = min(100, p.Hope+selected.HopeGain)

	// 特殊描述
	if selected.SanityGain > 0 {
		color.Green("理智值恢复了 %d 点", selected.SanityGain)
	} else if selected.SanityGain < 0 {
		color.Red("理智值下降了 %d 点", -selected.SanityGain)
	}
	
	if selected.HopeGain > 0 {
		color.Green("希望值提升了 %d 点", selected.HopeGain)
	}

	// 随机附加事件
	if rand.Intn(100) < 30 {
		color.HiBlack("\n（喝太快呛到了，损失￥10）")
		p.Spend(10)
	}

	utils.WaitForInput()
}