package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ferziza/pet_project/internal/db"
	"github.com/ferziza/pet_project/internal/repository"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(" .env файл не найден")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}

	ctx := context.Background()
	database, err := db.NewDatabase(ctx)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer database.Close()

	repo := repository.NewPlayerRepository(database)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Обработка нажатий на кнопки
		if update.CallbackQuery != nil {
			handleCallback(bot, update.CallbackQuery, repo, ctx)
			continue
		}

		// Обработка текстовых сообщений
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text
		messageID := update.Message.MessageID
		switch text {
		case "/start":

			SendMainMenu(bot, chatID)
		case "back_to_menu":
			editMsg := tgbotapi.NewEditMessageText(chatID, messageID, " Главное меню\n\nВыбери действие:")

			keyboard := GetMainMenuKeyboard()
			editMsg.ReplyMarkup = &keyboard

			bot.Send(editMsg)
		default:
			SendMainMenu(bot, chatID)
		}
	}
}
