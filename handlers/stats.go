package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/derekslenk/gomodels"
)

type stats struct {
	EpisodeCount        int `json:"EpisodeCount"`
	SpecialEpisodeCount int `json:"SpecialEpisodeCount"`
	AvgLength           int `json:"AvgLength"`
	AvgLengthSpecial    int `json:"AvgLengthSpecial"`
}

// StatsIndex returns default page for stats
func StatsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Stats Index Hit")
	stats := stats{
		EpisodeCount:        models.EpCount(),
		SpecialEpisodeCount: models.SpecialEpCount(),
		AvgLength:           models.AvgLength(),
		AvgLengthSpecial:    models.AvgLengthSpecial(),
	}
	json.NewEncoder(w).Encode(stats)
}
