package main

import "github.com/hajimehoshi/ebiten/v2"

func TryMovePlayer(g *Game) {
	players := g.WorldTags["players"]

	x := 0
	y := 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		pos.X += x
		pos.Y += y
	}
}
