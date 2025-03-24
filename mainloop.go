package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

func mainLoop() {
	player := Player{
		Sanity:      100,
		Hope:        100,
		ResumeCount: 0,
		Rejections:  0,
		IsGodMode:   false,
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
			gameOver()
			break
		}

		if player.Sanity <= 0 {
			color.HiMagenta("\n系统检测到精神崩溃...")
			time.Sleep(2 * time.Second)
			gameOver()
			break
		}

		showStatus(&player)

		// 动态选项系统（根据状态变化）
		options := []string{
			color.HiBlueString("[1] 海投简历（自我安慰版）"),
			color.HiBlueString("[2] 刷LeetCode（假装努力）"),
			color.HiBlueString("[3] 查看邮箱（直面现实）"),
			color.HiBlueString("[4] 喝奶茶（短暂多巴胺）"),
			color.HiRedString("[5] 打开BOSS直聘（地狱模式）"),
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
		input := strings.TrimSpace(getInput())

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
			drinkMilkTea(&player)
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
		default:
			color.Red("无效操作！你的犹豫消耗了时间...")
			player.Sanity -= 5
			time.Sleep(1 * time.Second)
		}

		// 随机事件触发
		if rand.Intn(100) > 80 {
			randomEvent(&player)
		}
	}
}

// 重写简历（自我欺骗）
func rewriteResume(p *Player) {
	clearScreen()
	
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
	waitForInput()
}

// 联系猎头（付费挨骂）
func headhunter(p *Player) {
	clearScreen()
	
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
	waitForInput()
}

// 随机事件系统
func randomEvent(p *Player) {
	events := []func(*Player){
		// 正能量事件（极少）
		func(p *Player) {
			if rand.Intn(100) > 95 { // 5%概率
				color.HiGreen("\n[系统] 收到意外赞美！")
				color.White("某HR在后台备注：")
				color.Green("'这人简历虽然菜但照片挺好看'")
				p.Hope += 20
			}
		},
		// 负能量事件
		func(p *Player) {
			color.Red("\n[紧急] 班级群通知：")
			color.White("你的室友 %s 拿到了%soffer",
				color.HiGreenString("张三"),
				color.HiYellowString("50万/年"))
			p.Hope -= 25
		},
		func(p *Player) {
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
	waitForInput()
}