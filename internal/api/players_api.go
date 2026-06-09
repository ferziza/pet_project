package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ferziza/pet_project/internal/dto"
)

func FetchAllPlayers(ctx context.Context) ([]dto.PlayerDTO, error) {
	var allPlayers []dto.PlayerDTO
	skip := 0
	take := 100

	for {
		url := fmt.Sprintf("https://ltd-hub.com/api/players?skip=%d&take=%d", skip, take)

		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var response dto.PlayersResponse
		json.Unmarshal(body, &response)

		allPlayers = append(allPlayers, response.Players...)

		if len(response.Players) < take {
			break
		}

		skip += take
		time.Sleep(100 * time.Millisecond)
	}

	return allPlayers, nil
}
