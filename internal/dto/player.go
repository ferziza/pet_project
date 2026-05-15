package dto

import "time"

type PlayerDTO struct {
	IrinaID     string    `json:"irianbotId"`
	Nickname    string    `json:"nickname"`
	WinsCount   int       `json:"winsCount"`
	LosesCount  int       `json:"losesCount"`
	LeavesCount int       `json:"leavesCount"`
	Rating      int       `json:"rating"`
	Rank        int       `json:"rank"`
	Winrate     float64   `json:"winrate"`
	Total       int       `json:"total"`
	BrokenCount int       `json:"brokenCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PlayersResponse struct {
	Players []PlayerDTO `json:"players"`
	Total   int         `json:"total"`
}
