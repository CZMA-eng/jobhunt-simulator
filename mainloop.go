package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"jobhunt/consumables"
	"jobhunt/players"
	"jobhunt/utils"
	"jobhunt/randomEvents"
	"github.com/fatih/color"
)

func mainLoop() {
	player := players.Player{
		Sanity:      100,
		Hope:        100,
		ResumeCount: 0,
		Rejections:  0,
		IsGodMode:   false,
		Money:       3000,
	}

	// 隐藏作弊码检测
	cheatCodes := map[string]func(){
		"/godmode": func() {
			player.IsGodMode = !player.IsGodMode
			color.HiMagenta("上帝模式已切换: %v", player.IsGodMode)
			time.Sleep(1 * time.Second)
		},
		"/suicide": func() {
			secretEnding(&player)
		},
		"/hope": func() {
			player.Hope = 100
			color.HiGreen("希望值回满！ (系统错误：这不应该发生)")
			time.Sleep(1 * time.Second)
		},
	}

	for {
		// 死亡检测
		if player.Hope <= 0 {
			color.HiRed("\n你的希望已归零...")
			time.Sleep(2 * time.Second)
			utils.GameOver(&player)
			break
		}

		if player.Sanity <= 0 {
			color.HiMagenta("\n系统检测到精神崩溃...")
			time.Sleep(2 * time.Second)
			utils.GameOver(&player)
			break
		}

		players.ShowStatus(&player)

		// 动态选项系统（根据状态变化）
		options := []string{
			color.HiBlueString("[1] 海投简历（自我安慰版）"),
			color.HiBlueString("[2] 刷LeetCode（假装努力）"),
			color.HiBlueString("[3] 查看邮箱（直面现实）"),
			color.HiBlueString("[4] 喝奶茶（短暂多巴胺）"),
			color.HiRedString("[5] 打开BOSS直聘（地狱模式）"),
			color.HiGreenString("[q] 退出游戏"),
		}

		// 特殊状态选项
		if player.Rejections > 10 {
			options = append(options, color.HiBlackString("[6] 重写简历（自我欺骗）"))
		}
		if player.ResumeCount > 50 {
			options = append(options, color.HiYellowString("[9] 联系猎头（付费挨骂）"))
		}

		fmt.Println("\n可选操作：")
		for _, opt := range options {
			fmt.Println(opt)
		}

		fmt.Print("\n> ")
		input := strings.TrimSpace(utils.GetInput())

		// 作弊码检测
		if action, exists := cheatCodes[input]; exists {
			action()
			continue
		}

		switch input {
		case "1":
			massApply(&player)
		case "2":
			leetCode(&player)
		case "3":
			checkEmail(&player)
		case "4":
			consumables.BuyDrink(&player)
		case "5":
			bossMode(&player)
		case "6":
			if player.Rejections > 10 {
				rewriteResume(&player)
			}
		case "9":
			if player.ResumeCount > 50 {
				headhunter(&player)
			}
		case "q":
			color.Red("心态崩了吗？ 现实可能也没好哪里去哦～ :)")
			fmt.Println()
			return
		default:
			color.Red("无效操作！你的犹豫消耗了时间...")
			player.Sanity -= 5
			time.Sleep(1 * time.Second)
		}

		// 随机事件触发
		if rand.Intn(100) > 80 {
			randomEvents.RandomEvent(&player)
		}
	}
}

// 重写简历（自我欺骗）
func rewriteResume(p *players.Player) {
	utils.ClearScreen()
	
	color.Cyan("第%d次重写简历...", p.Rejections)
	color.White("正在生成技术栈：")
	
	// 随机生成虚假技能
	techs := []string{
		"精通量子计算",
		"全栈（指前后端+运维+设计+产品）",
		"三年AI经验（实际调过API）",
		"诺贝尔奖提名（待确认）",
	}
	
	for i := 0; i < 3; i++ {
		fmt.Printf("- %s\n", color.YellowString(techs[rand.Intn(len(techs))]))
		time.Sleep(1 * time.Second)
	}
	
	color.Red("\n警告：检测到简历膨胀！")
	p.Sanity -= 20
	p.Hope += 10 // 虚假的希望
	utils.WaitForInput()
}

// 联系猎头（付费挨骂）
func headhunter(p *players.Player) {
	utils.ClearScreen()
	
	color.Red("▓▓▓ 高端人才顾问 ▓▓▓")
	color.White("服务费：998元/次")
	color.White("正在接入专业顾问...")
	fakeLoading(5)
	
	advices := []string{
		"你的要价是市场价的3倍",
		"30岁还投Junior岗位？",
		"建议转行送外卖",
		"你的简历像拼多多商品页",
	}
	
	color.Cyan("\n顾问诊断：")
	color.HiRed(advices[rand.Intn(len(advices))])
	
	p.Sanity -= 40
	p.Hope -= 30
	utils.WaitForInput()
}

