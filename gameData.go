package main


type GameData struct {
	ScreenWidth int
	ScreenHeight int
}



func NewGameData() GameData {
	g := GameData{
		ScreenWidth: 1280,
		ScreenHeight: 800,
	}
	return g
}
