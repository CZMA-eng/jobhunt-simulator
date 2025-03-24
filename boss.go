package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

var bossEvents = []func(*Player){
	hrGhosting,       // 已读不回
	lowballOffer,     // 恶意压价
	pyramidScheme,    // 传销式招聘
	ageDiscrimination,// 年龄歧视
}

func bossMode(p *Player) {
	clearScreen()
	
	color.Red("▓▓▓▓▓▓▓ BOSS直聘 - 地狱模式 ▓▓▓▓▓▓▓")
	fmt.Println()
	
	color.White("正在连接猎头网络...")
	fakeLoading(2)
	
	// 每次打开必定触发伤害
	p.ApplyDamage()
	p.ApplyDamage() // 双重打击
	
	// 随机触发一个黑暗事件
	event := bossEvents[rand.Intn(len(bossEvents))]
	event(p)
	
	waitForInput()
}

// 事件1: 已读不回
func hrGhosting(p *Player) {
	color.Cyan("\n【神秘HR】")
	color.White("查看你的简历")
	color.HiBlack("最后上线时间: 2秒前")
	
	// 模拟正在输入
	fmt.Print("\n对方正在输入...")
	time.Sleep(3 * time.Second)
	color.Red("（输入已停止）")
	
	p.Sanity -= 25
	color.HiMagenta("\n* 你的理智裂开了一道缝 *")
}

// 事件2: 恶意压价
func lowballOffer(p *Player) {
	color.Cyan("\n【创业公司CEO】")
	color.White("我们非常欣赏你的潜力！")
	color.White("但鉴于你：")
	color.HiBlack("✓ 缺乏996经验")
	color.HiBlack("✓ 不是图灵奖得主")
	color.HiRed("只能提供薪资: ￥3,000/月")
	
	fmt.Println()
	color.White("是否接受？ [y/n]")
	if getInput() == "y" {
		color.Red("\n签约成功！获得成就：")
		color.Red("██ 人矿认证 ██")
		p.Hope = 0
	} else {
		color.Red("\n对方撤回了一个offer")
	}
	p.Sanity -= 40
}

// 事件3: 传销式招聘
func pyramidScheme(p *Player) {
	color.Cyan("\n【财富自由导师】")
	color.White("加入我们的Web3.0生态矩阵！")
	color.HiGreen("✓ 无需经验")
	color.HiGreen("✓ 月入百万")
	color.HiGreen("✓ 发展下线奖励")
	
	fmt.Println()
	color.White("输入'我要发财'立即加入> ")
	if getInput() == "我要发财" {
		color.HiYellow("\n恭喜成为第114514号代理！")
		color.HiBlack("（你的通讯录已被上传）")
		p.IsGodMode = true // 开启无敌模式（黑色幽默）
	} else {
		color.Red("\n你失去了阶级跃迁的最后机会")
		p.Hope -= 50
	}
}

// 事件4: 年龄歧视（隐藏暴力机制）
func ageDiscrimination(p *Player) {
	color.Cyan("\n【系统提示】")
	color.White("检测到你的年龄可能超过")
	color.HiRed("25岁（程序员退休年龄）")
	
	fmt.Println()
	color.White("需要验证身份：")
	color.White("[1] 我是实习生")
	color.White("[2] 我出生时就会写代码")
	color.White("[3] 承认自己老了")
	
	switch getInput() {
	case "1":
		color.Red("\n实习经历过多，疑似跳槽倾向")
	case "2":
		color.Red("\n检测到早衰迹象")
	case "3":
		color.Red("\n年龄验证通过：不适合任何岗位") 
	}
	
	// 无论怎么选都会受伤
	p.Sanity -= 35
	p.Hope -= 20
	color.HiBlack("\n* 听到大脑碎裂的声音 *")
}

// 作弊码触发隐藏结局
func secretEnding(p *Player) {
	clearScreen()
	color.HiMagenta(`
███████████████████████
   系统错误：404 NOT FOUND
███████████████████████

检测到非法逃离求职系统的企图...
正在执行惩罚协议：

✓ 永久标记为"不稳定人员"
✓ 所有招聘平台拉黑
✓ 父母收到失败通知

（按下Ctrl+C 以确认人生失败）`)
	os.Exit(0)
}