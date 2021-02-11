package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct{}

type Position struct {
	X int
	Y int
}

func (p *Position) GetManhattanDistance(other *Position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))
	return int(xDist) + int(yDist)
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

type Monster struct {
	Name string
}
