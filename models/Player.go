package models

type Role int

const (
   Don Role = iota
   Mafia 
   Sheriff
   Civilian
)

type Player struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Role Role `json:"role"`
	Alive bool `json:"alive"`
	Warnings int `json:"warnings"`
	Eliminated bool `json:"eliminated"`
}