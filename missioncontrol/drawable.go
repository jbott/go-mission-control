package missioncontrol

import "github.com/veandco/go-sdl2/sdl"

type drawableInterface interface {
	Draw(*sdl.Renderer)
	Update()
}

type drawable struct {
	X            int32
	Y            int32
	W            int32
	H            int32
	border_color [4]uint8
	fill         bool
	fill_color   [4]uint8
}

func (d *drawable) Draw(r *sdl.Renderer) {
	// Fill
	if d.fill {
		r.SetDrawColor(
			d.fill_color[0],
			d.fill_color[1],
			d.fill_color[2],
			d.fill_color[3])
		r.FillRect(&sdl.Rect{d.X, d.Y, d.W, d.H})
	}

	// Draw bounds
	r.SetDrawColor(
		d.border_color[0],
		d.border_color[1],
		d.border_color[2],
		d.border_color[3])
	r.DrawRect(&sdl.Rect{d.X, d.Y, d.W, d.H})
}

func (d *drawable) Update() {}

func (d *drawable) SetBorderColor(r uint8, g uint8, b uint8, a uint8) {
	d.border_color[0] = r
	d.border_color[1] = g
	d.border_color[2] = b
	d.border_color[3] = a
}

func (d *drawable) SetFillColor(r uint8, g uint8, b uint8, a uint8) {
	d.fill = true
	d.fill_color[0] = r
	d.fill_color[1] = g
	d.fill_color[2] = b
	d.fill_color[3] = a
}
