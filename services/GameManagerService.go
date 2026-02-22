package services

import (
	"errors"
	"fmt"
	"main.go/constants"
	"main.go/models"
)

type GameManagerService interface {
	RegisterPlayer(name string, role constants.Role) (models.Player, error)
	StartGame(players []models.Player) error
	EliminatePlayer(playerName string) error
	RemovePlayer(playerName string) error
	GetAlivePlayers() []models.Player
	GetPlayerRole(playerName string) (constants.Role, error)
	EndGame(alicvePlayers []models.Player) (string, error)
}

type gameManager struct {
	players []models.Player
}

func NewGameManager() GameManagerService {
	return &gameManager{
		players: []models.Player{},
	}
}

// Implement the methods of GameManagerService interface here
func (gm *gameManager) RegisterPlayer(name string, role constants.Role) (models.Player, error) {
	newPlayer := models.Player{
		Id: len(gm.players) + 1,
		Name: name,
		Role: role,
		Alive: true,
		Warnings: 0,
		Eliminated: false,
	}

	gm.players = append(gm.players, newPlayer);

	return newPlayer, nil
}

func (gm *gameManager) StartGame(players []models.Player) error {	
	fmt.Printf("Starting game with %d players...\n", len(players))
	return nil
}

func (gm *gameManager) GetAlivePlayers() []models.Player {
	alivePlayers := []models.Player{}
	for _, p := range gm.players {
		if p.Alive {
			alivePlayers = append(alivePlayers, p)
		}
	}
	return alivePlayers
}

func (gm *gameManager) GetPlayerRole(playerName string) (constants.Role, error) {
	for _, p := range gm.players {
		if p.Name == playerName {
			return p.Role, nil
		}
	}
	return 0, errors.New("player not found: " + playerName)
}

func (gm *gameManager) RemovePlayer(playerName string) error {
	targetIndex := -1

    // 1. Manual loop through the SLICE of structs
    for i, p := range gm.players {
        if p.Name == playerName { // Access the field directly
            targetIndex = i
            break
        }
    }

    // 2. The Safety Check
    if targetIndex == -1 {
        return errors.New("player not found: " + playerName)
    }

    // 3. NOW - Apply the "Swap and Pop" we learned!
    gm.players[targetIndex] = gm.players[len(gm.players)-1] // Swap with the last element
	gm.players = gm.players[:len(gm.players)-1] // Pop the last element
    
    return nil
}

func (gm *gameManager) EliminatePlayer(playerName string) error {
	for index, p := range gm.players {
		if p.Name == playerName {
			gm.players[index].Alive = false
			gm.players[index].Eliminated = true
			return nil
		}
	}

	return errors.New("PLayer not found");
}

func (gm *gameManager) EndGame(alicvePlayers []models.Player) (string, error) {
	mafiaCount := 0
	civilianCount := 0

	for _, p := range alicvePlayers {
		switch p.Role {
        case constants.Mafia, constants.Don:
			mafiaCount++
		case constants.Civilian, constants.Sheriff:
			civilianCount++
		}
	}

	if mafiaCount == 0 {
		return string(constants.Civilians), nil
	} 
	
	return string(constants.Criminals), nil
	
}