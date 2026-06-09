package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendMainMenu(bot *tgbotapi.BotAPI, chatID int64) {
	button := tgbotapi.NewInlineKeyboardButtonData("Топ 10", "top_10")

	row := tgbotapi.NewInlineKeyboardRow(button)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	msg := tgbotapi.NewMessage(chatID, "Используйте кнопки для навигации")

	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}
func GetMainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	button := tgbotapi.NewInlineKeyboardButtonData("🏆 Топ 10 игроков", "top_10")
	row := tgbotapi.NewInlineKeyboardRow(button)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}
