package tg_client

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendMessage(token string, chatID int64, message string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, message)
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func HandleButtons(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	buttonData := update.CallbackQuery.Data
	callbackID := update.CallbackQuery.ID

	switch buttonData {
	case "button_1", "button_2", "button_3", "button_4":
		text := fmt.Sprintf("You pressed %s", buttonData)
		buttons := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Button 5", "button_5"),
			tgbotapi.NewInlineKeyboardButtonData("Back to menu", "back_to_menu"),
		}
		sendInlineKeyboard(bot, chatID, text, buttons)

	case "button_5":
		text := fmt.Sprintf("You pressed %s", buttonData)
		buttons := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Back to main menu", "back_to_first_menu"),
		}
		sendInlineKeyboard(bot, chatID, text, buttons)

	case "back_to_menu", "back_to_first_menu":
		text := fmt.Sprintf("You pressed %s", buttonData)
		buttons := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Button 1", "button_1"),
			tgbotapi.NewInlineKeyboardButtonData("Button 2", "button_2"),
			tgbotapi.NewInlineKeyboardButtonData("Button 3", "button_3"),
			tgbotapi.NewInlineKeyboardButtonData("Button 4", "button_4"),
		}
		sendInlineKeyboard(bot, chatID, text, buttons)

	case "button_6", "button_7":
		text := fmt.Sprintf("You pressed %s", buttonData)
		buttons := []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonData("Back to main menu", "back_to_first_menu"),
		}
		sendInlineKeyboard(bot, chatID, text, buttons)

	default:
		bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
			CallbackQueryID: update.CallbackQuery.ID,
			Text:            "Invalid button",
			ShowAlert:       true,
		})
	}

	bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackID, ""))
}

func sendInlineKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string, buttons []tgbotapi.InlineKeyboardButton) {
	message := tgbotapi.NewMessage(chatID, text)
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)
	bot.Send(message)
}
