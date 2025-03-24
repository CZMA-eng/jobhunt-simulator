package boss

import (
	"jobhunt/players"
	"jobhunt/utils"
	"time"
)

var resumePollutionLevel int

func CorruptResume(p *players.Player) {
	utils.ColorPrint(utils.ColorHiBlack, "\n[检测到简历正在被AI篡改...]")
	utils.ShowLoading(3, "loading")
	
	mods := []struct{
		text string
		pollution int
	}{
		{"添加虚假的NASA工作经历", 30},
		{"植入加密货币挖矿脚本", 50},
		{"生成舔狗式自我评价", 20},
	}
	
	for _, mod := range mods {
		utils.ColorPrint(utils.ColorHiYellow, "✓ %s", mod.text)
		resumePollutionLevel += mod.pollution
		time.Sleep(1 * time.Second)
	}
	
	if resumePollutionLevel >= 100 {
		unlockResumeHacks(p)
	}
}

func unlockResumeHacks(p *players.Player) {
	utils.ColorPrint(utils.ColorHiRed, "\n[简历完成赛博格化]")
	hacks := []string{
		"✓ 自动海投机器人",
		"✓ 假项目生成器",
		"✓ 面试AI替身系统",
	}
	utils.PrintBulletList(hacks, utils.ColorHiBlack)
	p.IsGodMode = true
}