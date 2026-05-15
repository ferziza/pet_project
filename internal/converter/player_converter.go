package converter

import (
	"time"

	"github.com/ferziza/pet_project/internal/dto"
	"github.com/ferziza/pet_project/internal/models"
)

func ToModel(p *dto.PlayerDTO) *models.Player {
	return &models.Player{

		IrinaID:      p.IrinaID,
		Nickname:     p.Nickname,
		WinsCount:    p.WinsCount,
		LossesCount:  p.LosesCount,
		LeavesCount:  p.LeavesCount,
		Rating:       p.Rating,
		RankPosition: p.Rank,
		WinRate:      p.Winrate,
		TotalGames:   p.Total,
		BrokenCount:  p.BrokenCount,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
