package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

func startInterview(p *Player) {
	clearScreen()
	color.HiMagenta("▓▓▓ 面试炼狱 ▓▓▓")
	fmt.Println()

	// 随机选择面试类型
	switch rand.Intn(3) {
	case 0:
		conductInterview("tech.md", p)
	case 1:
		conductInterview("hr.md", p)
	case 2:
		conductInterview("brainteaser.md", p)
	}
}

func conductInterview(file string, p *Player) {
	// 读取面试题
	questions, err := readInterviewQuestions(file)
	if err != nil {
		color.Red("面试官迷路了...")
		return
	}

	color.White("面试官 %s 正在审视你...", getInterviewerTitle(file))
	waitForInput()

	// 随机抽取3道题
	for i := 0; i < 3; i++ {
		q := questions[rand.Intn(len(questions))]
		color.HiCyan("\n问题 %d/%d:", i+1, 3)
		color.White(q)

		fmt.Print("\n你的回答 > ")
		getInput() // 无论回答什么都会被怼

		// 随机负面反馈
		feedbacks := []string{
			"面试官皱了皱眉",
			"HR在笔记本上画了个叉",
			"听到对面传来冷笑声",
			"这个问题你居然没准备？",
		}
		color.Red(feedbacks[rand.Intn(len(feedbacks))])
		p.Sanity -= 15
	}

	// 面试结果（10%概率通过）
	if rand.Intn(10) == 0 {
		color.HiGreen("\n奇迹发生！你通过了面试！")
		color.HiBlack("（接下来是6轮交叉面试）")
		p.Hope += 30
	} else {
		color.Red("\n感谢您的时间！")
		color.HiBlack("（已加入人才库-黑洞版）")
		p.Rejections++
	}
}

func readInterviewQuestions(file string) ([]string, error) {
	// 实际开发中从文件读取，这里简化示例
	switch file {
	case "tech.md":
		return []string{
			"请用O(1)时间复杂度给宇宙热寂加速",
			"解释React的虚拟DOM如何拯救世界和平",
			"如果让你设计一个能承受银河系级流量的系统...",
		}, nil
	case "hr.md":
		return []string{
			"你最大的缺点是什么？（必须说一个会被当成优点的缺点）",
			"如果CEO的狗不喜欢你怎么办？",
			"你愿意免费加班到公司上市吗？",
		}, nil
	default:
		return []string{
			"井盖为什么是圆的？",
			"如何用激光测地球到月球的距离？只能使用excel",
			"上海有多少个理发师？",
		}, nil
	}
}

func getInterviewerTitle(file string) string {
	titles := map[string]string{
		"tech.md":      "首席PUA工程师",
		"hr.md":       "人才优化专家",
		"brainteaser.md": "脑筋急转弯十段",
	}
	return titles[file]
}