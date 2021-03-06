package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("2133739093:AAHJasR6PQZa2BnQ5bFAHEo9hb1RfsOXWj8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		// if !update.Message.IsCommand() {
		// 	continue
		// }

		//Create a new MessageConfig. We don't have text yet, so we leave it empty
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message
		// switch update.Message.Command() {
		// case "help":
		// 	msg.Text = "I understand /sayhi and /status."
		// case "sayhi":
		// 	msg.Text = "Hi :)"
		// case "status":
		// 	msg.Text = "I'm ok ^^"
		// default:
		// 	msg.Text = "I'm not sure what that command entails."
		// }

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}
}
