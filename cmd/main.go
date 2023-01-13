package main

import (
	"fmt"
	"mutabornaya/pkg/database"
	"mutabornaya/pkg/mutaborec"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("BOT_TOKEN")

	sa := option.WithCredentialsFile("serviceAccounts.json")
	db, err := database.NewDatabase(sa)
	if err != nil {
		fmt.Fprintf(os.Stderr, "db error %v\n", err)
		os.Exit(1)
	}

	m, err := mutaborec.NewBot(db, token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bot error %v\n", err)
		os.Exit(1)
	}
	defer m.Close()

	m.Start()
	// bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// botUser, err := bot.GetMe()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// // Print Bot information
	// fmt.Printf("Bot user: %+v\n", botUser)

	// updates, _ := bot.UpdatesViaLongPulling(nil)

	// // Stop reviving updates from updates channel
	// defer bot.StopLongPulling()

	// // Loop through all updates when they came
	// for update := range updates {
	// 	fmt.Println(update.MyChatMember)
	// 	if update.Message != nil {
	// 		if update.Message.Sticker != nil {
	// 			// Retrieve chat ID
	// 			fmt.Println(update.Message)
	// 			chatID := update.Message.Chat.ID

	// 			// Call method sendMessage.
	// 			// Sends message to sender with same text (echo bot).
	// 			// (https://core.telegram.org/bots/api#sendmessage)
	// 			bot.SendSticker(tu.Sticker(
	// 				tu.ID(chatID),
	// 				tu.FileByID(update.Message.Sticker.FileID),
	// 			))
	// 			sentMessage, _ := bot.SendMessage(
	// 				tu.Message(
	// 					tu.ID(chatID),
	// 					update.Message.Text,
	// 				),
	// 			)

	// 			fmt.Printf("Sent Message: %v\n", sentMessage)
	// 			chat, _ := bot.GetChat(&telego.GetChatParams{
	// 				ChatID: telego.ChatID{
	// 					ID: chatID,
	// 				},
	// 			})
	// 			go schedule.Schedule(context.Background(), time.Minute*1, 0, func(t time.Time) {
	// 				fmt.Println(1111)
	// 				send(bot, chatID)
	// 			})
	// 			fmt.Printf("chat: %v\n\n", chat.ActiveUsernames)
	// 		}
	// 	}
	// }

	// _, err = bot.SendMessage(&telego.SendMessageParams{
	// 	ChatID: telego.ChatID{
	// 		ID: -897279868,
	// 	},
	// 	Text: "kek",
	// })
	// fmt.Println(err)
}

// func send(bot *telego.Bot, chatID int64) {
// 	bot.SendMessage(
// 		tu.Message(
// 			tu.ID(chatID),
// 			"xuy",
// 		),
// 	)
// }
