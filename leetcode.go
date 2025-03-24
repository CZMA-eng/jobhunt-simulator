package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

var leetcodeQuestions = []struct {
	Problem string
	Choices []string
}{
	{
		"用O(1)时间复杂度实现时间倒流",
		[]string{
			"发明时光机",
			"修改服务器时间",
			"在简历上谎称会这个",
		},
	},
	{
		"反转黑洞的物质流向（需要分布式解决方案）",
		[]string{
			"使用Kafka做消息队列",
			"用Go routine处理奇点",
			"重启宇宙",
		},
	},
}

func leetCode(p *Player) {
	clearScreen()
	
	q := leetcodeQuestions[rand.Intn(len(leetcodeQuestions))]
	color.Cyan(q.Problem + "\n")
	
	for i, choice := range q.Choices {
		fmt.Printf("[%d] %s\n", i+1, color.WhiteString(choice))
	}
	
	// 假装有选择
	getInput()
	
	color.Red("\n运行结果：超出时间限制")
	color.HiBlack("（你的解题思路已被ChatGPT取代）")
	p.ApplyDamage()
	waitForInput()
}