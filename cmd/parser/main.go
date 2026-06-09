package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ferziza/pet_project/internal/api"
	"github.com/ferziza/pet_project/internal/converter"
	"github.com/ferziza/pet_project/internal/db"
	"github.com/ferziza/pet_project/internal/dto"
	"github.com/ferziza/pet_project/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️ .env файл не найден, используем системные переменные")
	}

	ctx := context.Background()

	database, err := db.NewDatabase(ctx)
	if err != nil {
		fmt.Printf(" Ошибка подключения к БД: %v\n", err)
		os.Exit(1)
	}
	defer database.Close()

	fmt.Println("Подключились к бд")

	repo := repository.NewPlayerRepository(database)

	for {
		ctxTimeout, cancel := context.WithTimeout(ctx, 1*time.Minute)
		players, err := api.FetchAllPlayers(ctxTimeout)
		cancel()
		if err != nil {
			fmt.Printf(" Ошибка API: %v\n", err)
			time.Sleep(1 * time.Minute)
			continue
		}

		var wg sync.WaitGroup

		for _, p := range players {
			wg.Add(1)
			go func(player dto.PlayerDTO) {
				defer wg.Done()
				if err := repo.Upsert(ctx, converter.ToModel(&player)); err != nil {
					fmt.Printf(" Ошибка сохранения %s: %v\n", player.Nickname, err)
				}
			}(p)
		}
		wg.Wait()

		time.Sleep(1 * time.Minute)
	}
}
