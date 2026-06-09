package repository

import (
	"context"
	"fmt"

	"github.com/ferziza/pet_project/internal/db"
	"github.com/ferziza/pet_project/internal/models"
	"github.com/jackc/pgx/v5"
)

type PlayerRepository struct {
	db *db.Database
}

func NewPlayerRepository(database *db.Database) *PlayerRepository {
	return &PlayerRepository{db: database}
}

func (r *PlayerRepository) Upsert(ctx context.Context, player *models.Player) error {
	query := `
        INSERT INTO players (
            irina_id, nickname, wins_count, losses_count, leaves_count,
            rating, rank_position, winrate, total_games, broken_count,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        ON CONFLICT (irina_id) DO UPDATE SET
            nickname = EXCLUDED.nickname,
            wins_count = EXCLUDED.wins_count,
            losses_count = EXCLUDED.losses_count,
            leaves_count = EXCLUDED.leaves_count,
            rating = EXCLUDED.rating,
            rank_position = EXCLUDED.rank_position,
            winrate = EXCLUDED.winrate,
            total_games = EXCLUDED.total_games,
            broken_count = EXCLUDED.broken_count,
            updated_at = EXCLUDED.updated_at
   			 `

	_, err := r.db.Pool.Exec(ctx, query,
		player.IrinaID,
		player.Nickname,
		player.WinsCount,
		player.LossesCount,
		player.LeavesCount,
		player.Rating,
		player.RankPosition,
		player.WinRate,
		player.TotalGames,
		player.BrokenCount,
		player.CreatedAt,
		player.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("ошибка сохранения игрока %s: %w", player.Nickname, err)
	}

	return nil
}

// получает топ N игроков по рейтингу
func (r *PlayerRepository) GetTopByRating(ctx context.Context, limit int) ([]models.Player, error) {
	query := `SELECT * FROM players ORDER BY rating DESC LIMIT $1`

	rows, err := r.db.Pool.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByPos[models.Player])
}
