package data 

import (
    "main.go/models"
    "main.go/constants"
)

var MockPlayersData = []models.Player{
    {Name: "Alice", Role: constants.Don},
    {Name: "Bob", Role: constants.Civilian},
    {Name: "Charlie", Role: constants.Mafia},
    {Name: "Diana", Role: constants.Civilian},
    {Name: "Inna", Role: constants.Civilian},
    {Name: "Ruben", Role: constants.Civilian},
    {Name: "Eve", Role: constants.Sheriff},
}