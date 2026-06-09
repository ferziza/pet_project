package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ferziza/pet_project/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, repo *repository.PlayerRepository, ctx context.Context) {
	chatID := message.Chat.ID
	text := message.Text

	switch text {
	case "/start":
		SendMainMenu(bot, chatID)

	default:
		msg := tgbotapi.NewMessage(chatID, "Используй кнопки из меню")
		bot.Send(msg)
		SendMainMenu(bot, chatID)
	}
}
func handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, repo *repository.PlayerRepository, ctx context.Context) {
	bot.Send(tgbotapi.NewCallback(callback.ID, ""))

	chatID := callback.Message.Chat.ID
	messageID := callback.Message.MessageID
	callbackData := callback.Data

	switch callbackData {
	case "top_10":
		sendTop10Players(bot, chatID, messageID, repo, ctx)

	case "back_to_menu":

		editMsg := tgbotapi.NewEditMessageText(chatID, messageID, "Главное меню\n\nВыбери действие:")

		keyboard := GetMainMenuKeyboard()
		editMsg.ReplyMarkup = &keyboard

		bot.Send(editMsg)

	default:
		SendMainMenu(bot, chatID)
	}
}
func sendTop10Players(bot *tgbotapi.BotAPI, chatID int64, messageID int, repo *repository.PlayerRepository, ctx context.Context) {
	players, err := repo.GetTopByRating(ctx, 10)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка получения данных")
		bot.Send(msg)
		return
	}

	if len(players) == 0 {
		msg := tgbotapi.NewMessage(chatID, "Данных пока нет")
		bot.Send(msg)
		return
	}

	var responseText strings.Builder
	responseText.WriteString("🏆 *ТОП 10 ИГРОКОВ ПО РЕЙТИНГУ*\n\n")

	for i, player := range players {

		responseText.WriteString(fmt.Sprintf("%d. *%s* - %d очков\n", i+1, player.Nickname, player.Rating))
		responseText.WriteString(fmt.Sprintf("   Побед: %d | Поражений: %d | Винрейт: %.1f%%\n\n",
			player.WinsCount, player.LossesCount, player.WinRate))
	}

	editMsg := tgbotapi.NewEditMessageText(chatID, messageID, responseText.String())
	editMsg.ParseMode = "Markdown"

	backButton := tgbotapi.NewInlineKeyboardButtonData("◀️ Назад в меню", "back_to_menu")
	backRow := tgbotapi.NewInlineKeyboardRow(backButton)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(backRow)
	editMsg.ReplyMarkup = &keyboard

	bot.Send(editMsg)
}
