package models

import "main.go/constants"

type Player struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Role constants.Role `json:"role"`
	Alive bool `json:"alive"`
	Warnings int `json:"warnings"`
	Eliminated bool `json:"eliminated"`
}