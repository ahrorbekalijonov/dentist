package models

//Standart Error
type StandartError struct {
	Error string 
}

//Error
type Error struct {
	Error StandartError
}