package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)


func ShowOpening() {
	ClearScreen()

	// 第一阶段：赛博朋克风格加载
	color.HiMagenta("初始化求职痛苦系统...")
	showCyberLoading(3)

	// 第二阶段：破碎ASCII艺术字
	ClearScreen()
	fig := figure.NewColorFigure("JOB HUNT 2025", "larry3d", "red", true)
	fig.Print()

	// 第三阶段：动态就业市场报告
	fmt.Println()
	showMarketReport()

	// 第四阶段：无厘头免责声明
	showDisclaimer()

	WaitForInput()
}

// ========== 酷炫组件 ==========

func showCyberLoading(seconds int) {
    colors := []color.Attribute{
        color.FgHiMagenta,
        color.FgHiBlue,
        color.FgHiCyan,
    }
    frames := []string{
        "[■□□□□□□□□□]",
        "[■■■□□□□□□□]",
        "[■■■■■□□□□□]",
        "[■■■■■■■□□□]",
        "[■■■■■■■■■□]",
        "[■■■■■■■■■■]",
    }

    for i := 0; i < len(frames); i++ {
        c := colors[i%len(colors)]
        // 修复百分比计算：(i+1)*100/len(frames)
        percent := (i+1)*100/len(frames)
        color.New(c).Printf("\r%s 加载中... %d%%", frames[i], percent)
        time.Sleep(time.Duration(seconds)*time.Second/time.Duration(len(frames)))
    }
    fmt.Println()
}

func showMarketReport() {
	reports := []string{
		"■ 就业市场AI分析报告 ■",
		"  - 你的专业热度: " + color.HiRedString("已冷冻") + " ❄️",
		"  - 岗位竞争比: " + color.HiYellowString("1:114514") + " 💩",
		"  - 平均薪资: " + color.HiBlackString("3k") + " (包含老板画饼价值2.5k)",
		"  - 996指数: " + strings.Repeat("🔥", 10),
		"  - 你的竞争力: " + color.HiGreenString("NULL"),
	}

	for _, line := range reports {
		// 打字机效果
		for _, c := range line {
			fmt.Print(string(c))
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println()
	}
}

func showDisclaimer() {
	fmt.Println()
	color.HiRed("⚠️ 免责声明 ⚠️")
	disclaimers := []string{
		"本游戏可能引起以下症状:",
		color.HiYellowString("• 突然查看招聘软件强迫症"),
		color.HiYellowString("• 听见'赋能'就呕吐"),
		color.HiYellowString("• 梦见LeetCode题目"),
		color.HiYellowString("• 对HR产生PTSD"),
		"",
		"继续游戏即表示你同意:",
		color.HiMagentaString("• 交出你的头发"),
		color.HiMagentaString("• 献祭三年寿命"),
		color.HiMagentaString("• 接受福报洗礼"),
	}

	for _, line := range disclaimers {
		fmt.Print("💀 ")
		for _, c := range line {
			fmt.Print(string(c))
			time.Sleep(30 * time.Millisecond)
		}
		fmt.Println()
	}
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func GameOver(p GameOverData) {
    ClearScreen()

    // 随机死亡音效
    deathSounds := []string{"💀", "☠️", "👻", "🤡", "💩"}
    deathEmoji := deathSounds[rand.Intn(len(deathSounds))]

    // 生成结局
    ending := generateEnding(
        p.GetSanity(),
        p.GetHope(),
        p.GetMoney(),
        p.GetResumeCount(),
        p.GetRejections(),
        p.GetGhostedCount(),
    )

    // 打印结局
    color.HiRed("\n G A M E   O V E R ")
    fmt.Println(deathEmoji + ending + deathEmoji)
    fmt.Println()

    // 显示统计数据
    showStatistics(p)

    // 隐藏结局检测
    if p.GetResumeCount() >= 100 && p.GetMoney() <= 0 {
        showHiddenEnding()
    }

    fmt.Println()
    color.HiBlack("按回车键回到现实世界...")
    WaitForInputWithPrompt(color.HiBlackString("按回车键回到现实世界..."))
    os.Exit(0)
}

// ========== 内部私有函数 ==========
func WaitForInputWithPrompt(prompt string) {
    fmt.Println()
    fmt.Print(prompt)
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func generateEnding(sanity, hope, money, resumes, rejects, ghosts int) string {
    switch {
    case sanity <= 0 && hope <= 0:
        return fmt.Sprintf(`
%s 终极崩溃 %s
  你同时获得了：
  ✓ 精神分裂症 %s
  ✓ 抑郁症 %s
  ✓ 求职 PTSD %s
  
  解锁成就：【%s】`,
            color.HiRedString("▄︻デ══━"),
            color.HiRedString("━══デ︻▄"),
            color.HiMagentaString("MAX"),
            color.HiBlueString("MAX"),
            color.HiYellowString("MAX"),
            color.HiRedString("人间不值得"))

    case rejects >= 100:
        return fmt.Sprintf(`
%s 专业陪跑员 %s
  简历投递数：%s
  收到拒信：%s
  默拒次数：%s
  
  特别成就：【%s】`,
            color.HiBlueString("▬▬ι═══════"),
            color.HiBlueString("═══════ι▬▬"),
            color.HiYellowString("%d", resumes),
            color.HiRedString("%d", rejects),
            color.HiBlackString("%d", ghosts),
            color.HiMagentaString("HR 数据库黑名单"))

//     case money <= 0:
//         return fmt.Sprintf(`
// %s 破产宣言 %s
//   最后余额：%s
//   简历打印负债：%s
//   奶茶续命花费：%s
  
//   获得称号：【%s】`,
//             color.HiYellowString("≪"),
//             color.HiYellowString("≫"),
//             color.HiRedString("￥%d", money),
//             color.HiRedString("￥998"),
//             color.HiRedString("￥%d", 3000-money),
//             color.HiBlackString("赛博乞丐"))

    default:
        endings := []string{
            fmt.Sprintf(`
%s 三和大神 %s
  解锁技能：
  ✓ 日结工资 %s
  ✓ 网吧生存 %s
  ✓ 社保规避 %s
  
  最终状态：【%s】`,
                color.HiGreenString("⟩⟩⟩"),
                color.HiGreenString("⟨⟨⟨"),
                color.HiYellowString("Lv.MAX"),
                color.HiBlueString("Lv.MAX"),
                color.HiMagentaString("Lv.MAX"),
                color.HiBlackString("社会性死亡")),

            fmt.Sprintf(`
%s 赛博精神病 %s
  症状包括：
  ✓ 看见JD就呕吐 %s
  ✓ 听见"赋能"抽搐 %s
  ✓ 梦见笔试题目 %s
  
  临床诊断：【%s】`,
                color.HiMagentaString("✧･ﾟ"),
                color.HiMagentaString("･ﾟ✧"),
                color.HiRedString("✔"),
                color.HiRedString("✔"),
                color.HiRedString("✔"),
                color.HiYellowString("晚期求职综合征")),

            fmt.Sprintf(`
%s 佛系咸鱼 %s
  最终属性：
  %s 理智：%d
  %s 希望：%d
  %s 负债：%d
  
  领悟真谛：【%s】`,
                color.HiBlueString("∵"),
                color.HiBlueString("∴"),
                color.HiCyanString("▸"),
                sanity,
                color.HiYellowString("▸"),
                hope,
                color.HiRedString("▸"),
                3000-money,
                color.HiBlackString("躺平才是终极答案")),
        }
        return endings[rand.Intn(len(endings))]
    }
}

func showStatistics(p GameOverData) {
    color.HiBlack("══════ 求职统计 ══════")
    fmt.Printf(" %s 简历投递: %d\n", color.HiBlueString("✉"), p.GetResumeCount())
    fmt.Printf(" %s 明确拒绝: %d\n", color.HiRedString("✖"), p.GetRejections())
    fmt.Printf(" %s 默拒次数: %d\n", color.HiBlackString("⊙"), p.GetGhostedCount())
    fmt.Printf(" %s 最后财产: %s\n", color.HiYellowString("💰"), color.HiRedString("￥%d", p.GetMoney()))
    fmt.Printf(" %s 精神创伤: %s\n", color.HiMagentaString("💊"), getTraumaLevel(p.GetSanity()))
}

func showHiddenEnding() {
    fmt.Println()
    color.HiYellow("隐藏结局解锁：")
    color.HiRed("  [专业炮灰] 你的简历已成为HR的垃圾邮件过滤器训练数据")
}

func getTraumaLevel(sanity int) string {
    switch {
    case sanity <= 0:
        return color.HiRedString("永久性损伤") + " 💀"
    case sanity <= 20:
        return color.HiRedString("严重创伤") + " 🤕"
    case sanity <= 50:
        return color.HiYellowString("中度抑郁") + " 😵"
    default:
        return color.HiGreenString("暂时性崩溃") + " 🥴"
    }
}