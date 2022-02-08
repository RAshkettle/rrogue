package main

//GameData holds the values for the size of elements within the game
type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
	UIHeight     int
}

//NewGameData creates a fully populated GameData Struct.
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 60,
		TileWidth:    16,
		TileHeight:   16,
		UIHeight:     10,
	}

	return g
}
