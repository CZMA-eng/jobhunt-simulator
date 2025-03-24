package boss

import (
	"math/rand"
	"jobhunt/players"
	"jobhunt/utils"
)

type AIInterviewer struct {
	Title    string
	Style    func(string) string
	Effect   func(*players.Player)
	Color    *utils.ColorPrinter
}

var aiInterviewers = []AIInterviewer{
	{
		"PUA大师",
		func(s string) string { return "🤔 " + s },
		func(p *players.Player) { 
			p.ModifySanity(-20)
			utils.ColorPrint(utils.ColorHiMagenta, "你的存在价值受到质疑") 
		},
		utils.NewColorPrinter(utils.ColorHiMagenta),
	},
	{
		"福报传教士",
		func(s string) string { return "🙏 " + s },
		func(p *players.Player) { 
			p.Hope -= 30
			utils.ColorPrint(utils.ColorHiYellow, "感受到福报的召唤") 
		},
		utils.NewColorPrinter(utils.ColorHiYellow),
	},
}

func StartAIInterview(p *players.Player) {
	ai := aiInterviewers[rand.Intn(len(aiInterviewers))]
	
	ai.Color.Printf("\n【AI面试官-%s】\n", ai.Title)
	
	questions := []string{
		"如果公司要求你付费上班，你会怎么选择？",
		"用三个字形容你作为工具人的觉悟",
		"你愿意为工作放弃哪些基本人权？",
	}
	
	for _, q := range questions {
		ai.Color.Print(ai.Style(q))
		utils.WaitForInput()
		ai.Effect(p)
	}
	
	utils.ColorPrint(utils.ColorHiBlack, "\n[AI面试官对你的评分已加入区块链]")
}