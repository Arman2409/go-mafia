package main

import (
	"fmt"

	"main.go/data"
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
	fmt.Println("Alive Players:")
	for _, p := range alivePlayers {
		fmt.Printf("- %s\n", p.Name) // Use %s for strings
	}
}
