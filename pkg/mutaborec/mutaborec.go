package mutaborec

import (
	"fmt"
	"mutabornaya/pkg/database"

	"github.com/mymmrac/telego"
)

type Mutaborec struct {
	bot *telego.Bot
	db  *database.Database
}

func NewBot(db *database.Database, botToken string) (*Mutaborec, error) {
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}

	m := &Mutaborec{
		db:  db,
		bot: bot,
	}

	return m, nil
}

func (m *Mutaborec) Start() {
	updates, _ := m.bot.UpdatesViaLongPulling(nil)

	for update := range updates {
		switch {
		case update.MyChatMember != nil:
			status := update.MyChatMember.NewChatMember.MemberStatus()
			if status == "member" {
				fmt.Println("join")
			}
			if status == "left" {
				fmt.Println("left")
			}
		case update.Message != nil:
		}
	}
}

func (m *Mutaborec) Close() {
	m.bot.StopLongPulling()
	m.db.Close()
}
