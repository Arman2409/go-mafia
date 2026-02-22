package constants

type Side string

const (
    Criminals Side = "Criminals"
    Civilians Side = "Civilians"
)

type Role int

const (
   Don Role = iota
   Mafia 
   Sheriff
   Civilian
)
