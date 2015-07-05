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
		SetDrawColor(r, d.fill_color)
		r.FillRect(&sdl.Rect{d.X, d.Y, d.W, d.H})
	}

	// Draw bounds
	SetDrawColor(r, d.border_color)
	r.DrawRect(&sdl.Rect{d.X, d.Y, d.W, d.H})
}

func (d *drawable) Update() {}

func (d *drawable) SetBorderColor(r uint8, g uint8, b uint8, a uint8) {
	SetColor(&d.border_color, r, g, b, a)
}

func (d *drawable) SetFillColor(r uint8, g uint8, b uint8, a uint8) {
	d.fill = true
	SetColor(&d.fill_color, r, g, b, a)
}
