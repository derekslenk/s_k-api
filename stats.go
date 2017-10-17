package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		EpisodeCount:        EpCount(),
		SpecialEpisodeCount: SpecialEpCount(),
		AvgLength:           AvgLength(),
		AvgLengthSpecial:    AvgLengthSpecial(),
	}
	json.NewEncoder(w).Encode(stats)
}
