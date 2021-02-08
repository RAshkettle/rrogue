package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		pos := result.Components[position].(*Position)
		img := result.Components[renderable].(*Renderable).Image

		if level.PlayerVisible.IsVisible(pos.X, pos.Y) {
			index := level.GetIndexFromXY(pos.X, pos.Y)
			tile := level.Tiles[index]
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(img, op)
		}

	}
}
