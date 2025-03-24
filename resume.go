package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func massApply(p *Player) {
    clearScreen()
    count := rand.Intn(5) + 3
    p.ResumeCount += count

    color.Cyan("投递 %d 份简历...", count)
    fakeLoading(2)

    newOpportunities := 0
    for i := 0; i < count; i++ {
        fate := rand.Intn(100)
        switch {
        case fate < 60: // 60%默拒
            p.GhostedCount++
        case fate < 90: // 30%待处理拒绝
            p.PendingReplies++
        default: // 10%面试机会
            p.PendingReplies++
            newOpportunities++
        }
    }

    // 显示汇总结果
    color.White("\n投递结果：")
    fmt.Printf("%s %d份石沉大海\n", color.HiBlackString("✉"), count-p.PendingReplies)
    fmt.Printf("%s %d份待处理\n", color.HiYellowString("→"), p.PendingReplies)
    if newOpportunities > 0 {
        color.HiGreen("🌟 获得 %d 个潜在面试机会！", newOpportunities)
        color.HiBlack("（请查看邮箱等待通知）")
    }

    p.ApplyDamage()
    waitForInput()
}
