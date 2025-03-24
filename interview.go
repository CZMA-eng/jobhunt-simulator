package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

// 面试状态记录
type InterviewProgress struct {
	CurrentRound    int
	MaxRounds       int
	PassedTech      bool
	PassedWhiteboard bool
	PassedHR        bool
	AskedQuestions  map[string]bool // 记录已问问题
}

func NewInterview() *InterviewProgress {
	return &InterviewProgress{
		CurrentRound: 1,
		MaxRounds:    rand.Intn(4) + 1, // 随机1-4轮
		AskedQuestions: make(map[string]bool),
	}
}

func startInterview(p *Player) {
	clearScreen()
	progress := NewInterview()
	color.HiMagenta("▓▓▓ 面试炼狱 ▓▓▓")
	color.HiBlack("（共%d轮面试）", progress.MaxRounds)
	
	for progress.CurrentRound <= progress.MaxRounds {
		if !conductRound(p, progress) {
			// 本轮未通过
			color.Red("\n很遗憾，你在第%d轮被淘汰", progress.CurrentRound)
			p.Rejections++
			waitForInput()
			return
		}
		
		progress.CurrentRound++
		if progress.CurrentRound <= progress.MaxRounds {
			color.HiGreen("\n进入第%d轮面试...", progress.CurrentRound)
			time.Sleep(2 * time.Second)
		}
	}
	
	// 所有轮次通过
	color.HiGreen("\n奇迹发生！你通过了所有面试！")
	color.HiBlack("（HR说还要等老板最终审批）")
	p.Hope += 50
	waitForInput()
}

func conductRound(p *Player, progress *InterviewProgress) bool {
	clearScreen()
	
	var roundType string
	var questions []string
	
	switch {
	case !progress.PassedTech:
		roundType = "技术面"
		questions = loadQuestions("tech/round"+fmt.Sprint(progress.CurrentRound)+".md")
	case !progress.PassedWhiteboard:
		roundType = "白板编程"
		questions = loadQuestions("tech/system_design.md")
	case !progress.PassedHR:
		roundType = "HR面"
		questions = loadQuestions("hr/questions.md")
	default:
		roundType = "终极面试"
		questions = loadQuestions("brainteaser/questions.md")
	}
	
	color.HiCyan("=== 第%d轮: %s ===", progress.CurrentRound, roundType)
	color.White("面试官: %s", getInterviewerTitle(roundType))
	fmt.Println()
	
	// 随机选择3个未问过的问题
	selected := selectUniqueQuestions(questions, 3, progress.AskedQuestions)
	if len(selected) == 0 {
		color.Red("面试官找不到问题了...")
		return true // 默认通过
	}
	
	// 回答环节
	correctNeeded := rand.Intn(2) + 1 // 需要答对1-2题
	correctCount := 0
	
	for i, q := range selected {
		color.HiCyan("\n问题 %d/%d:", i+1, len(selected))
		color.White(q)
		
		fmt.Print("\n你的回答 > ")
		getInput()
		
		// 随机判断是否正确
		if rand.Intn(100) < 60 { // 60%概率答对
			correctCount++
			color.HiGreen("✓ 面试官微微点头")
		} else {
			feedbacks := []string{
				"面试官摇了摇头",
				"听到对面传来叹息声",
				"这个问题答得不太行",
				"面试官对你的专业产生了质疑",
				"你感觉自己甚至不如国际学院学生",
			}
			color.Red(feedbacks[rand.Intn(len(feedbacks))])
		}
		
		p.Sanity -= 10
		if p.Sanity <= 0 {
			color.Red("\n你的精神崩溃了...")
			return false
		}
	}
	
	// 判断是否通过本轮
	if correctCount >= correctNeeded {
		// 更新进度
		switch roundType {
		case "技术面":
			progress.PassedTech = true
		case "白板编程":
			progress.PassedWhiteboard = true
		case "HR面":
			progress.PassedHR = true
		}
		return true
	}
	return false
}

// 从文件加载问题
func loadQuestions(path string) []string {
	fullPath := filepath.Join("interviews", path)
	file, err := os.Open(fullPath)
	if err != nil {
		return []string{"默认问题：请解释宇宙的终极答案"}
	}
	defer file.Close()
	
	var questions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			questions = append(questions, line)
		}
	}
	
	if len(questions) == 0 {
		return []string{"错误：题库为空"}
	}
	return questions
}

// 选择未问过的问题
func selectUniqueQuestions(all []string, count int, asked map[string]bool) []string {
	// 先找出所有未问过的问题
	var available []string
	for _, q := range all {
		if !asked[q] {
			available = append(available, q)
		}
	}
	
	// 如果可用问题不足，重置记录
	if len(available) < count {
		for k := range asked {
			delete(asked, k)
		}
		available = all
	}
	
	// 随机选择
	rand.Shuffle(len(available), func(i, j int) {
		available[i], available[j] = available[j], available[i]
	})
	
	selected := available[:min(count, len(available))]
	for _, q := range selected {
		asked[q] = true
	}
	
	return selected
}

func getInterviewerTitle(roundType string) string {
	titles := map[string]string{
		"技术面":     "首席PUA工程师",
		"白板编程":    "系统架构虐待狂",
		"HR面":     "人才优化专家",
		"终极面试":    "CEO亲信",
	}
	return titles[roundType]
}