package main

import (
	"fmt"
	"sync"
	"main.go/data"
	"main.go/models"
	"main.go/services"
)

func main() {
	mockPlayers := data.MockPlayersData
	gameManagerService := services.NewGameManager()

	for _, player := range mockPlayers {
		fmt.Printf("Player: %s, Role: %d\n", player.Name, player.Role)
		gameManagerService.RegisterPlayer(player.Name, player.Role)
	}

	gameManagerService.StartGame(mockPlayers)

	gameManagerService.GetPlayerRole("Bob")
	gameManagerService.EliminatePlayer("Charlie")
	gameManagerService.EliminatePlayer("Diana")

	alivePlayers := gameManagerService.GetAlivePlayers()
	winner, err := gameManagerService.EndGame(alivePlayers)

	if err != nil {
		fmt.Printf("Error determining winner: %s\n", err.Error())
	}

	statsManager := services.NewStatsManager(winner)

	var wg sync.WaitGroup

	for _, player := range mockPlayers {
	   wg.Add(1)
	   go func(p models.Player) {
		   defer wg.Done()
		   statsManager.CalculatePlayerStats(&p)
	   }(player)
	}

	wg.Wait()
	
	statsManager.GetOverallStats()
}
