package missioncontrol

import "github.com/veandco/go-sdl2/sdl"

type widget struct {
	X int32
	Y int32
	W int32
	H int32
}

func NewWidget(x int32, y int32) *widget {
	return &widget{x, y, 0, 0}
}

func (w *widget) Draw(r *sdl.Renderer) {
	// Draw bounds
	points := []sdl.Point{{w.X, w.Y}, {w.X + w.W, w.Y}, {w.X + w.W, w.Y + w.H}, {w.X, w.Y + w.H}, {w.X, w.Y}}
	r.SetDrawColor(255, 255, 255, 255)
	r.DrawLines(points)
}
