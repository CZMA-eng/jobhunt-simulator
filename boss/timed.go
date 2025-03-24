package boss

import (
	"fmt"
	"time"
	"jobhunt/players"
	"jobhunt/utils"
)

func StartTimedChallenge(p *players.Player, question string, timeout int) bool {
	fmt.Printf("\n%s (%d秒倒计时开始)\n", question, timeout)
	
	result := make(chan bool)
	defer close(result)
	
	// 倒计时动画
	go func() {
		for i := timeout; i > 0; i-- {
			fmt.Printf("\r剩余时间: %d ", i)
			time.Sleep(1 * time.Second)
		}
		result <- false
	}()
	
	// 输入监听
	go func() {
		utils.GetInput() // 阻塞等待输入
		result <- true
	}()
	
	if success := <-result; success {
		utils.ColorPrint(utils.ColorHiGreen, "提交成功！")
		return true
	}
	
	utils.ColorPrint(utils.ColorHiRed, "时间耗尽！")
	p.ModifySanity(-40)
	return false
}