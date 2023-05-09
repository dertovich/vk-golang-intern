package main

import (
	"log"
	"vk-golang-intern/tg_client"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	config, err := tg_client.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tg_client.CreateBot(config)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = config.DebugMode

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go tg_client.HandleButtons(bot, update)
		} else if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, I'm just a bot created for an internship at VK.")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Button 1"),
					tgbotapi.NewKeyboardButton("Button 2"),
				),
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Button 3"),
					tgbotapi.NewKeyboardButton("Button 4"),
				),
			)
			bot.Send(msg)
		}
	}
}
