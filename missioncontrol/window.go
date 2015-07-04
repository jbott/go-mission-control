package missioncontrol

import "github.com/veandco/go-sdl2/sdl"

type window struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Widgets []*widget
}

func NewWindow(x int32, y int32) *window {
	w := new(window)
	w.X = x
	w.Y = y
	return w
}

func (w *window) AddWidget(widget *widget) {
	w.Widgets = append(w.Widgets, widget)
}

func (w *window) Draw(r *sdl.Renderer) {
	// Draw bounds
	//points := []sdl.Point{{w.X, w.Y}, {w.X + w.W, w.Y}, {w.X + w.W, w.Y + w.H}, {w.X, w.Y + w.H}, {w.X, w.Y}}
	r.SetDrawColor(255, 255, 0, 0)
	r.DrawRect(&sdl.Rect{w.X, w.Y, w.W, w.H})
	r.DrawLine(200, 200, 400, 400)

	// Draw Widgets
	for _, widget := range w.Widgets {
		widget.Draw(r)
	}
}
