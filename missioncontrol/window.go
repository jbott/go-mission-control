package missioncontrol

import "github.com/veandco/go-sdl2/sdl"

type window struct {
	drawable
	Widgets []*widget
}

func (mc *MissionControl) NewWindow(x int32, y int32) *window {
	w := new(window)
	w.X = x
	w.Y = y
	w.border_color = mc.default_border_color
	return w
}

func (w *window) AddWidget(widget *widget) {
	w.Widgets = append(w.Widgets, widget)
}

func (w *window) Draw(r *sdl.Renderer) {
	w.drawable.Draw(r)

	// Draw Widgets
	for _, widget := range w.Widgets {
		widget.Draw(r)
	}
}
