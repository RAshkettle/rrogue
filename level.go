package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

//MapTile is a single Tile on a given level
type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

//NewLevel creates a new game level in a dungeon.
func NewLevel() Level {
	l := Level{}

	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()

	return l
}

//DrawLevel draws the level onto the screen.
func (level *Level) DrawLevel(screen *ebiten.Image) {
	gd := NewGameData()

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {

			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}

}

//GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
//This coordinate is logical tiles, not pixels.
func (level *Level) GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

//GenerateLevelTiles creates a new Dungeon Level Map.
func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	tiles := level.createTiles()
	level.Tiles = tiles

	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		x := GetDiceRoll(gd.ScreenWidth-w-1) - 1
		y := GetDiceRoll(gd.ScreenHeight-h-1) - 1
		new_room := NewRect(x, y, w, h)

		okToAdd := true

		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}
		if okToAdd {
			level.createRoom(new_room)
			if len(level.Rooms) == 0 {

			}

			level.Rooms = append(level.Rooms, new_room)

		}
	}

}

//createTiles creates a map of all walls as a baseline for carving out a level.
func (level *Level) createTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, gd.ScreenHeight*gd.ScreenWidth)
	index := 0
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}
			tile := MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TileHeight,
				Blocked: true,
				Image:   wall,
			}
			tiles[index] = tile

		}
	}
	return tiles
}

func (level *Level) createRoom(room Rect) {
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}
