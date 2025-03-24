package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"jobhunt/players"
	"jobhunt/utils"
	"github.com/fatih/color"
)

var rejectionMessages = []string{
	color.HiBlackString("【自动回复】感谢投递"),
	fmt.Sprintf("%s 您的%s不匹配我们的要求", 
		color.HiWhiteString("很遗憾"),
		color.HiRedString("人生经历")),
	color.HiBlueString("【人才库】") + "已将您的简历存入黑洞",
	fmt.Sprintf("%s：月薪3k的全栈岗位考虑吗？", 
		color.HiYellowString("诈骗邮件")),
	fmt.Sprintf("%s\n%s", 
		color.HiWhiteString("面试邀请："),
		color.HiGreenString("诚聘厕所所长（可转正）")),
}

func checkEmail(p *players.Player) {
    utils.ClearScreen()
    
    if p.PendingReplies == 0 {
        color.White("收件箱空空如也...")
        color.HiBlack("(但你知道还有%d份简历被默拒)", p.GhostedCount)
        utils.WaitForInput()
        return
    }
    
    color.White("正在连接邮件服务器...")
    fakeLoading(3)
    
    // 确保至少处理1个面试邀请（如果有）
    processCount := 1
    if p.PendingReplies > 1 {
        processCount += rand.Intn(2) // 额外处理0-2封
    }
    if processCount > p.PendingReplies {
        processCount = p.PendingReplies
    }
    
    color.White("\n收到 %d 封新邮件：", processCount)
    
    for i := 0; i < processCount; i++ {
        p.PendingReplies--
        
        // 如果有面试机会优先处理（修改概率权重）
        if rand.Intn(100) < 30 || p.PendingReplies == 0 { 
            // 30%概率或最后一封必定是面试邀请
            handleInterviewInvite(p)
        } else {
            p.Rejections++
            color.HiRed("✖ " + rejectionMessages[rand.Intn(len(rejectionMessages))])
        }
    }
    
    showEmailStats(p)
    p.ApplyDamage()
    utils.WaitForInput()
}

func handleInterviewInvite(p *players.Player) {
    color.HiGreen("\n✓ 面试邀请：%s", getInterviewType())
    color.White("输入 'accept' 接受挑战 > ")
    
    if strings.ToLower(utils.GetInput()) == "accept" {
        startInterview(p)
    } else {
        color.Red("已自动拒绝机会")
        p.Hope -= 15
        // 拒绝后转成普通拒信
        p.Rejections++
        color.HiRed("✖ " + rejectionMessages[rand.Intn(2)]) 
    }
}

func showEmailStats(p *players.Player) {
    fmt.Printf("\n%s 待处理：%d | 默拒：%d | 总拒绝：%d\n",
        color.HiBlackString("统计："),
        p.PendingReplies,
        p.GhostedCount,
        p.Rejections)
}

func getInterviewType() string {
    types := []string{
        "技术面（手写操作系统）",
        "HR面（压力测试）",
        "CTO终极拷问",
        "白板编程（禁用编译器）",
    }
    return types[rand.Intn(len(types))]
}

func fakeLoading(seconds int) {
	fmt.Print("[")
	for i := 0; i < 20; i++ {
		fmt.Print(" ")
	}
	fmt.Print("]\r[")
	
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(seconds)*50 * time.Millisecond)
		fmt.Print("█")
	}
}