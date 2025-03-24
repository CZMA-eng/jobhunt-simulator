package randomEvents

import (
	"jobhunt/players"
	"jobhunt/utils"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// 地狱级随机事件系统
func RandomEvent(p *players.Player) {
	events := []func(*players.Player){
		// 表面正能量实则诛心事件
		func(p *players.Player) {
			color.HiGreen("\n[喜讯] 收到猎头电话！")
			color.White("高薪岗位：%s", color.HiYellowString("元宇宙厕所所长"))
			color.White("要求：")
			color.HiBlack("✓ 十年相关经验")
			color.HiBlack("✓ 精通量子力学擦屁股法")
			color.HiBlack("✓ 接受007工作制")
			p.Sanity -= 20
			utils.WaitForInput()
		},

		// 连环打击事件组
		func(p *players.Player) {
			color.HiRed("\n[致命连击]")
			color.Red("早上8:00  收到拒信：%s", color.HiBlackString("已加入人才库（黑洞版）"))
			color.Red("上午10:00 房贷催款通知：剩余￥%d", p.Money/2)
			color.Red("下午3:00  BOSS直聘消息：%s", 
				color.HiYellowString("已读不回"))
			color.Red("晚上11:00 母亲来电：%s", 
				color.HiWhiteString("邻居家孩子都当总监了"))
			p.Hope -= 40
			p.Sanity -= 30
			p.Money -= 500 // 被迫给家里打钱
		},

		// 黑色幽默诈骗事件
		func(p *players.Player) {
			color.HiMagenta("\n[官方邮件] 简历优化服务！")
			color.White("只需￥998 马上拿到Offer！")
			color.White("付款后获得：")
			color.HiBlack("✓ 自动编造大厂经历")
			color.HiBlack("✓ 伪造GitHub提交记录")
			color.HiBlack("✓ 代写离职证明生成器")
			
			if utils.Confirm("是否付款？") {
				p.Money -= 998
				color.HiRed("\n付款成功！获得：")
				color.Red("✓ 简历被AI标记为失信人员")
				color.Red("✓ 进入HR黑名单数据库")
				p.Rejections += 10
			} else {
				color.HiBlack("\n系统自动回复：穷逼不配找工作")
				p.Hope -= 15
			}
		},

		// 时间杀手事件
		func(p *players.Player) {
			color.HiYellow("\n[系统提示] 发现新功能！")
			color.White("完成性格测试可获得内推机会")
			color.White("正在加载300道测试题...")
			utils.ShowLoading(5, "浪费生命进度")
			color.Red("测试结果：%s", 
				color.HiBlackString("不适合在地球生存"))
			p.Sanity -= 25
			p.Hope -= 20
		},

		// 魔幻现实主义事件
		func(p *players.Player) {
			color.HiMagenta("\n[元宇宙HR通知]")
			color.White("您申请的%s岗位：", 
				color.HiYellowString("Web4.0架构师"))
			color.White("需先完成：")
			color.HiBlack("✓ 购买NFT面试通行证（￥999）")
			color.HiBlack("✓ 参加元宇宙笔试（48小时）")
			color.HiBlack("✓ 提供前世工作证明")
			p.Hope -= rand.Intn(30)
		},

		// 物理打击事件
		func(p *players.Player) {
			color.HiRed("\n[现实暴击]")
			color.White("合租室友的%s：", 
				color.HiGreenString("狗"))
			color.Red("✓ 在你的笔记本上撒尿")
			color.Red("✓ 吃掉最后一包泡面")
			color.Red("✓ 简历成为撕咬玩具")
			color.HiBlack("（维修费用扣除￥250）")
			p.Money -= 250
			p.Sanity -= 35
		},

		// 次元壁破裂事件 
		func(p *players.Player) {
			color.HiBlue("\n[穿越通知]")
			color.White("您已进入996平行宇宙：")
			color.HiBlack("✓ 时间流速 ×3")
			color.HiBlack("✓ 所有截止时间提前")
			color.HiBlack("✓ HR获得读心术能力")
			p.Sanity = max(0, p.Sanity-rand.Intn(50))
		},

		// 人工智能暴击
		func(p *players.Player) {
			color.HiWhite("\n[AI评估报告]")
			color.White("根据大数据分析：")
			color.HiRed("✓ 你的简历被机器人评为E级")
			color.HiRed("✓ 岗位匹配度：0.01%")
			color.HiRed("✓ 预计失业时长：∞")
			color.White("优化建议：%s", 
				color.HiBlackString("重新投胎"))
			p.Hope = max(0, p.Hope-30)
		},

		// 终极哲学打击
		func(p *players.Player) {
			color.HiYellow("\n[宇宙真相]")
			color.White("突然意识到：")
			color.HiBlack("✓ 你的专业是屠龙之术")
			color.HiBlack("✓ 所有岗位都是幻觉")
			color.HiBlack("✓ 人生是HR的模拟游戏")
			p.Sanity = 0 // 直接清空理智值
		},
	}

	// 概率黑箱（让玩家感受命运的恶意）
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100)
	switch {
	case r < 5:  // 5% 表面正能量
		events[0](p)
	case r < 40: // 35% 普通打击
		events[1+rand.Intn(3)](p)
	case r < 70: // 30% 重度打击 
		events[4+rand.Intn(3)](p)
	default:     // 30% 致命打击
		events[7+rand.Intn(3)](p)
	}

	// 确保数值不会溢出
	p.Hope = clamp(p.Hope, 0, 100)
	p.Sanity = clamp(p.Sanity, 0, 100)
	p.Money = max(p.Money, 0) // 允许负债吗？更现实！
}

// 辅助函数
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}