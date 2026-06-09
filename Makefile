include .env
export

# Запуск всего проекта (БД + парсер + бот)
up:
	@echo "🚀 Запуск PostgreSQL, парсера и бота..."
	@docker-compose up -d
	@echo "✅ Проект запущен"
	@echo "📊 Для просмотра логов: make logs"

down:
	@echo "🛑 Остановка PostgreSQL, парсера и бота..."
	@docker-compose down
	@echo "✅ Проект остановлен"

# Перезапуск всего проекта
restart: down up

# Просмотр логов (всех сервисов)
logs:
	@docker-compose logs -f

# Просмотр логов только парсера
logs-parser:
	@docker-compose logs -f parser

# Просмотр логов только бота
logs-bot:
	@docker-compose logs -f bot

# Статус контейнеров
status:
	@docker-compose ps
	@echo ""
	@docker-compose exec postgres psql -U postgres -d waric -c "SELECT COUNT(*) as players_count FROM players;" 2>/dev/null || echo "БД не доступна"


# Обновление парсера (после изменений кода)
update-parser:
	@echo "🔄 Пересборка и перезапуск парсера..."
	@docker-compose up -d --build parser
	@echo "✅ Парсер обновлен"
	@docker-compose logs -f parser

# Обновление бота (после изменений кода)
update-bot:
	@echo "🔄 Пересборка и перезапуск бота..."
	@docker-compose up -d --build bot
	@echo "✅ Бот обновлен"
	@docker-compose logs -f bot

# Обновление всего
update: update-parser update-bot

# Перезапуск только бота
restart-bot:
	@echo "🔄 Перезапуск бота..."
	@docker-compose restart bot
	@echo "✅ Бот перезапущен"
	@docker-compose logs -f bot

# Очистка всего (удалить контейнеры и тома)
clean:
	@echo "🧹 Очистка контейнеров и томов..."
	@docker-compose down -v
	@echo "✅ Очистка завершена"