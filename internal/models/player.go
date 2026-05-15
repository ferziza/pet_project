package models

import "time"

type Player struct {
	ID           int64     `db:"id"`
	IrinaID      string    `db:"irina_id"`
	Nickname     string    `db:"nickname"`
	WinsCount    int       `db:"wins_count"`
	LossesCount  int       `db:"losses_count"`
	LeavesCount  int       `db:"leaves_count"`
	Rating       int       `db:"rating"`
	RankPosition int       `db:"rank_position"`
	WinRate      float64   `db:"winrate"`
	TotalGames   int       `db:"total_games"`
	BrokenCount  int       `db:"broken_count"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
