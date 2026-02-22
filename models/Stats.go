package models 

type PlayerStats struct {
	PlayerName string `json:"player_name"`
	GamesPlayed int `json:"games_played"`
	GamesWon int `json:"games_won"`
	Killed int `json:"killed"`
}