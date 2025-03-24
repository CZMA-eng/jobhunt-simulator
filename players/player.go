package players

import (
	"fmt"
	"github.com/fatih/color"
	"jobhunt/utils"
	"math/rand"
)

type Player struct {
	Sanity         int
	Hope           int
	ResumeCount    int
	Rejections     int
	GhostedCount   int // 新增：记录被默拒的数量
	PendingReplies int // 新增：待处理的回复数量
	IsGodMode      bool
	Money          int // 生活费
}

func (p *Player) ModifySanity(i int) {
	p.Sanity = i
}

// 确保Player实现GameOverData接口
var _ utils.GameOverData = (*Player)(nil)

func (p *Player) GetSanity() int       { return p.Sanity }
func (p *Player) GetHope() int         { return p.Hope }
func (p *Player) GetMoney() int        { return p.Money }
func (p *Player) GetResumeCount() int  { return p.ResumeCount }
func (p *Player) GetRejections() int   { return p.Rejections }
func (p *Player) GetGhostedCount() int { return p.GhostedCount }

func (p *Player) Spend(amount int) bool {
	if p.IsGodMode {
		return true
	}

	if p.Money >= amount {
		p.Money -= amount
		return true
	}

	color.Red("余额不足！当前剩余：￥%d", p.Money)
	return false
}

func (p *Player) ApplyDamage() {
	if p.IsGodMode {
		return
	}

	p.Sanity -= rand.Intn(20)
	p.Hope -= rand.Intn(15)
	if p.Sanity < 0 {
		p.Sanity = 0
	}
	if p.Hope < 0 {
		p.Hope = 0
	}
}

func ShowStatus(p *Player) {
	utils.ClearScreen()

	// 创建颜色对象
	sanityColor := color.New(color.FgHiCyan)
	hopeColor := color.New(color.FgHiYellow)

	// 绘制状态条
	sanityBar := generateBar(p.Sanity, "█", sanityColor)
	hopeBar := generateBar(p.Hope, "▓", hopeColor)

	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf(`
%s 当前状态 %s
  理智值: %s 
  希望值: %s 
  生活费: %s
  已投简历: %s%d份
  收到拒信: %s%d份
`,
		cyan("▛▜"), cyan("▙▟"),
		sanityBar,
		hopeBar, color.HiYellowString("￥%d", p.Money), // 格式化货币,
		yellow("≋"), p.ResumeCount,
		color.HiRedString("✖"), p.Rejections,
	)
	fmt.Printf("  简历状态: %s%d待处理 %s%d默拒 %s%d拒绝\n",
		color.HiYellowString("→"), p.PendingReplies,
		color.HiBlackString("⊙"), p.GhostedCount,
		color.HiRedString("✖"), p.Rejections)
}

// 修改函数签名，接收 *color.Color
func generateBar(percentage int, block string, c *color.Color) string {
	full := int(float64(percentage) / 100.0 * 20)
	bar := ""
	for i := 0; i < 20; i++ {
		if i < full {
			bar += c.Sprint(block)
		} else {
			bar += " "
		}
	}
	return bar + fmt.Sprintf(" %d%%", percentage)
}
