package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
//This coordinate is logical tiles, not pixels.
func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}

	return g
}

func CreateTiles() []MapTile {
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
