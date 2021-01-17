package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//Level holds the tile information for a complete dungeon level.
type Level struct {
	Tiles []MapTile
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
	tiles := l.CreateTiles()
	l.Tiles = tiles

	return l
}

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

func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	tiles := make([]MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {

			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
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

				tiles = append(tiles, tile)

			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TileHeight,
					Blocked: false,
					Image:   floor,
				}

				tiles = append(tiles, tile)
			}

		}
	}
	return tiles
}
