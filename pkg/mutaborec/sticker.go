package mutaborec

import (
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Unique sticker ids
const (
	PusherTony = "AgADBCEAAnMXYUk"
)

func (m *Mutaborec) handleSticker(msg *telego.Message) {
	chatID := msg.Chat.ID
	s := msg.Sticker

	if s.FileUniqueID == PusherTony {
		_, err := m.bot.SendSticker(tu.Sticker(
			tu.ID(chatID),
			tu.FileByID(s.FileID),
		))

		fmt.Println(err)
	}
}
