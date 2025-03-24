package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func drinkMilkTea(p *Player) {
	clearScreen()
	
	types := []string{
		"芋泥波波",
		"黑糖珍珠",
		"芝士葡萄",
		"香菜柠檬（地狱特调）",
	}
	choice := types[rand.Intn(len(types))]
	
	color.Magenta("你猛吸了一口 %s 奶茶...", choice)
	time.Sleep(2 * time.Second)
	
	if choice == "香菜柠檬（地狱特调）" {
		color.Red("\n这味道太可怕了！")
		p.Sanity -= 30
	} else {
		color.Green("\n短暂的多巴胺分泌...")
		p.Hope += 15
		p.Sanity += 10
		// 防止溢出
		if p.Hope > 100 {
			p.Hope = 100
		}
		if p.Sanity > 100 {
			p.Sanity = 100
		}
	}
	
	color.HiBlack("\n（糖分加速了你的焦虑）")
	waitForInput()
}