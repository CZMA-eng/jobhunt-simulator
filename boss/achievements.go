package boss

import (
	"jobhunt/players"
	"jobhunt/utils"
)

type Achievement struct {
	Title      string
	Check      func(*players.Player) bool
	UnlockText string
}

var achievements = []Achievement{
	{
		"福报先驱",
		func(p *players.Player) bool { return p.Rejections >= 996 },
		"你的简历被刻在资本纪念碑上",
	},
	{
		"人形爬虫", 
		func(p *players.Player) bool { return p.ResumeCount >= 5000 },
		"触发招聘网站反爬虫机制",
	},
	{
		"赛博乞丐",
		func(p *players.Player) bool { return p.Money <= 0 },
		"解锁街头乞讨模式：在CBD唱《感恩的心》",
	},
}

func checkAchievements(p *players.Player) {
	for _, a := range achievements {
		if a.Check(p) {
			showUnlockEffect(a)
			// TODO: 保存成就状态
		}
	}
}

func showUnlockEffect(a Achievement) {
	utils.ColorPrint(utils.ColorHiMagenta, "\n成就解锁：%s", a.Title)
	utils.ColorPrint(utils.ColorHiBlack, "➔ %s", a.UnlockText)
	utils.PlaySoundEffect("achievement_unlock")
}