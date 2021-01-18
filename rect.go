package main

type Rect struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func NewRect(x int, y int, width int, height int) Rect {
	return Rect{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

func (r *Rect) Center() (int, int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2

	return centerX, centerY

}

func (r *Rect) Intersect(other Rect) bool {
	return (r.X1 <= other.X2 && r.X2 >= other.X1 && r.Y1 <= other.Y1 && r.Y2 >= other.Y1)
}
