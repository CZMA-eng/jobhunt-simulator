package boss

import (
	"fmt"
	"jobhunt/players"
	"jobhunt/utils"
	"math/rand"
	"os"

	"github.com/fatih/color"
)

var Events = []func(*players.Player){
	hrGhosting,
	lowballOffer,
	pyramidScheme,
	ageDiscrimination,
	offerTrap,
	codeTest,
}

func triggerRandomEvent(p *players.Player) {
	// 权重随机：30%基础事件，20%特殊事件
	switch rand.Intn(100) {
	case 0-69:   Events[rand.Intn(4)](p)      // 基础事件
	case 70-89:  Events[4](p)                // Offer陷阱
	default:     Events[5](p)                // 代码测试
	}
}

// 已读不回事件
func hrGhosting(p *players.Player) {
	utils.ColorPrint(utils.ColorHiCyan, "\n【神秘HR】")
	utils.TypewriterEffect("查看你的简历...", 100)
	utils.ColorPrint(utils.ColorHiBlack, "最后上线时间: 2秒前")
	
	utils.ShowDynamicInputIndicator(3) // 模拟输入中动画
	utils.ColorPrint(utils.ColorHiRed, "（输入已停止）")
	
	p.ModifySanity(-15)
	utils.ColorPrint(utils.ColorHiMagenta, "* 理智出现裂痕 *")
}

// 恶意压价事件
func lowballOffer(p *players.Player) {
	utils.ColorPrint(utils.ColorHiCyan, "\n【创业公司CEO】")
	utils.ColorPrint(utils.ColorWhite, "我们非常欣赏你的潜力！")
	utils.ColorPrint(utils.ColorWhite, "但鉴于以下原因：")
	
	reasons := []string{
		"✓ 缺乏007工作制经验",
		"✓ 未获得图灵奖",
		"✓ 发际线不符合要求",
	}
	utils.PrintBulletList(reasons, utils.ColorHiBlack)
	
	utils.ColorPrint(utils.ColorHiRed, "最终报价：￥3,000/月")
	
	if utils.Confirm("接受这份福报吗？") {
		utils.ColorPrint(utils.ColorHiRed, "签约成功！")
		utils.ColorPrint(utils.ColorHiBlack, "（获得成就：人矿认证）")
		p.Hope = 0
	} else {
		utils.ColorPrint(utils.ColorHiRed, "对方撤回了offer")
	}
	p.ModifySanity(-20)
}

// 事件3: 传销式招聘
func pyramidScheme(p *players.Player) {
	color.Cyan("\n【财富自由导师】")
	color.White("加入我们的Web3.0生态矩阵！")
	color.HiGreen("✓ 无需经验")
	color.HiGreen("✓ 月入百万")
	color.HiGreen("✓ 发展下线奖励")
	
	fmt.Println()
	color.White("输入'我要发财'立即加入> ")
	if utils.GetInput() == "我要发财" {
		color.HiYellow("\n恭喜成为第114514号代理！")
		color.HiBlack("（你的通讯录已被上传）")
		p.IsGodMode = true // 开启无敌模式（黑色幽默）
	} else {
		color.Red("\n你失去了阶级跃迁的最后机会")
		p.Hope -= 50
	}
}

// 事件4: 年龄歧视（隐藏暴力机制）
func ageDiscrimination(p *players.Player) {
	color.Cyan("\n【系统提示】")
	color.White("检测到你的年龄可能超过")
	color.HiRed("25岁（程序员退休年龄）")
	
	fmt.Println()
	color.White("需要验证身份：")
	color.White("[1] 我是实习生")
	color.White("[2] 我出生时就会写代码")
	color.White("[3] 承认自己老了")
	
	switch utils.GetInput() {
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
func SecretEnding(p *players.Player) {
	utils.ClearScreen()
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

// Offer陷阱：表面高薪实则卖身契
func offerTrap(p *players.Player) {
	utils.ColorPrint(utils.ColorHiGreen, "\n★ ★ ★ 顶级Offer ★ ★ ★")
	utils.TypewriterEffect("【元宇宙福报架构师】", 50)
	utils.ColorPrint(utils.ColorHiYellow, "年薪: ￥1,500,000")

	// 用极小字号显示条款
	utils.ColorPrint(utils.ColorHiBlack, "────────────────────────")
	utils.ColorPrint(utils.ColorHiBlack, "条款6.3.9: 自愿放弃所有法定节假日")
	utils.ColorPrint(utils.ColorHiBlack, "条款8.8.8: 公司有权根据福报指数调整薪资")
	utils.ColorPrint(utils.ColorHiBlack, "条款9.9.9: 离职需支付工位空气使用费")
	utils.ColorPrint(utils.ColorHiBlack, "────────────────────────")

	utils.ColorPrint(utils.ColorHiCyan, "\nCEO视频邀请中...")
	utils.ShowDynamicInputIndicator(2)
	
	utils.ColorPrint(utils.ColorHiRed, "（突然弹出人脸识别窗口）")
	utils.ColorPrint(utils.ColorWhite, "请眨眼以确认接受所有条款>")
	
	if utils.Confirm("立即签约成为人上人吗？") {
		utils.ColorPrint(utils.ColorHiMagenta, "\n签约成功！获得：")
		utils.ColorPrint(utils.ColorHiBlack, "✓ 终身奋斗者证书")
		utils.ColorPrint(utils.ColorHiBlack, "✓ 公司logo纹身贴纸")
		utils.ColorPrint(utils.ColorHiRed, "（Hope值已清零）")
		p.Hope = 0
	} else {
		utils.ColorPrint(utils.ColorHiRed, "\n系统检测到缺乏奉献精神")
		utils.ColorPrint(utils.ColorHiBlack, "（所有公司对你的好感度下降）")
		p.ModifySanity(-15)
		p.Hope += 20 // 保持清醒反而增加希望
	}
}

// 代码测试：永远无法通过的谜题
func codeTest(p *players.Player) {
	utils.ColorPrint(utils.ColorHiBlue, "\n【天才少年CTO】")
	utils.ColorPrint(utils.ColorWhite, "请用太空语言(SpaceLang)实现：")
	utils.ColorPrint(utils.ColorHiCyan, "量子波动排序算法")

	// 代码编辑器动画
	utils.ColorPrint(utils.ColorHiBlack, "┌───────────────────────┐")
	utils.ColorPrint(utils.ColorHiBlack, "│ ░░░░░░░░░░░ 正在加载AI │")
	utils.ColorPrint(utils.ColorHiBlack, "└───────────────────────┘")
	utils.ShowDynamicInputIndicator(3)

	// 无论怎么写都出错
	problems := []string{
		"错误: 缺少表情包注释",
		"警告: 头发密度不达标",
		"异常: 未检测到996基因",
	}
	problem := problems[rand.Intn(len(problems))]  // 完全随机选择
	utils.ColorPrint(utils.ColorHiRed, problem)

	// 限时生死题
	utils.ColorPrint(utils.ColorHiRed, "\n[30秒内回答]")
	utils.ColorPrint(utils.ColorWhite, "如何用0字节内存解决")
	utils.ColorPrint(utils.ColorWhite, "旅行商NP-Hard问题？")

	choices := []string{
		"1. 质问出题人是否清醒",
		"2. 写个AI用魔法打败魔法",
		"3. 自曝曾用Windows XP",
	}
	utils.PrintBulletList(choices, utils.ColorHiWhite)

	switch utils.GetInput() {
	case "1":
		utils.ColorPrint(utils.ColorHiRed, "\n系统判定：缺乏成长型思维")
		utils.ColorPrint(utils.ColorHiBlack, "（已自动转发微博吐槽内容）")
	case "2":
		utils.ColorPrint(utils.ColorHiRed, "\nAI生成代码：")
		utils.ColorPrint(utils.ColorHiBlack, "fmt.Println(🚀) // 量子完成")
		utils.ColorPrint(utils.ColorHiRed, "（出现神秘语法错误）")
	case "3":
		utils.ColorPrint(utils.ColorHiRed, "\n年龄验证失败：XP是上古时代")
	default:
		utils.ColorPrint(utils.ColorHiRed, "\n思考超时，已自动提交空白代码")
	}

	// 隐藏暴力机制
	p.ModifySanity(-20)
	utils.ColorPrint(utils.ColorHiMagenta, "* 听见大脑蓝屏的声音 *")
	
	// 5%概率触发隐藏剧情
	if rand.Intn(100) < 5 {
		utils.ColorPrint(utils.ColorHiYellow, "\n（你偷偷在代码注释里写了SOS）")
		utils.ColorPrint(utils.ColorHiBlack, "三天后收到神秘邮件...")
		SecretEnding(p)
	}
}