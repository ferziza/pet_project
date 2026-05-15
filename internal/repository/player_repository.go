package repository

import (
	"context"

	"github.com/ferziza/pet_project/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PlayerRepository interface {
	Upsert(ctx context.Context, player *models.Player) error

	GetByID(ctx context.Context, id int64) (*models.Player, error)

	GetByIrinaID(ctx context.Context, irinaID string) (*models.Player, error)

	GetByNickname(ctx context.Context, nickname string) (*models.Player, error)

	GetTopPlayers(ctx context.Context, limit int) ([]models.Player, error)

	GetTopPlayersWithPagination(ctx context.Context, offset, limit int) ([]models.Player, error)

	GetTotalCount(ctx context.Context) (int, error)

	Delete(ctx context.Context, id int64) error
}

type playerRepository struct {
	db *pgxpool.Pool
}

func NewPlayerRepository(db *pgxpool.Pool) PlayerRepository {
	return &playerRepository{db: db}
}
