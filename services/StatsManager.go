package services

import (
	"errors"
	"fmt"

	"main.go/constants"
	"main.go/models"
)

type StatsManagerService interface {
	GetPlayerStats(playerName string) (models.PlayerStats, error)
	CalculatePlayerStats(player *models.Player) error
	GetOverallStats()
}

type StatsManager struct {
	stats []models.PlayerStats
	winner string
	eliminated int
}

func NewStatsManager(winner string) StatsManagerService {
	return &StatsManager{
		stats: []models.PlayerStats{},
		winner: winner,
	}
}

func (sm *StatsManager) GetPlayerStats(playerName string) (models.PlayerStats, error) {
	for _, stat := range sm.stats {
		if stat.PlayerName == playerName {
			return stat, nil
		}
	}

	return models.PlayerStats{}, errors.New("player stats not found")
}

func (sm *StatsManager) CalculatePlayerStats(player *models.Player) error {
    if(player == nil) {
		return errors.New("player is nil")
	}

	stats := models.PlayerStats{
		PlayerName: player.Name,
		GamesPlayed: 0, // This should be calculated based on game history
		GamesWon: 0,    // This should be calculated based on game history
		Killed: 0,      // This should be calculated based on game history
	}

	if(player.Eliminated) {
		sm.eliminated += 1
	}

	if(player.Role == constants.Mafia || player.Role == constants.Don) && sm.winner == string(constants.Criminals) {
		stats.GamesWon += 1
	} else if (player.Role == constants.Civilian || player.Role == constants.Sheriff) && sm.winner == string(constants.Civilians) {
		stats.GamesWon += 1
	}

	sm.stats = append(sm.stats, stats)

	return nil
}

func (sm *StatsManager) GetOverallStats() {
	for _, stat := range sm.stats {
		fmt.Printf("Player: %s, Games Played: %d, Games Won: %d, Killed: %d\n", stat.PlayerName, stat.GamesPlayed, stat.GamesWon, stat.Killed)
	}
}