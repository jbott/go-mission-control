package missioncontrol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type MissionControl struct {
	Display              int
	drawables            []drawableInterface
	sdl_window           *sdl.Window
	renderer             *sdl.Renderer
	background_color     [4]uint8
	default_border_color [4]uint8
}

func Init() *MissionControl {
	mc := new(MissionControl)

	var err error

	// Create a new window with a default size
	mc.sdl_window, err = sdl.CreateWindow("Mission Control", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	// Set window to fullscreen
	var disp_rect sdl.Rect
	sdl.GetDisplayBounds(mc.Display, &disp_rect)
	mc.sdl_window.SetPosition(int(disp_rect.X), int(disp_rect.Y))
	mc.sdl_window.SetSize(int(disp_rect.W), int(disp_rect.H))
	mc.sdl_window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	mc.renderer, err = sdl.CreateRenderer(mc.sdl_window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	mc.renderer.Clear()

	mc.background_color[3] = 255

	return mc
}

func (mc *MissionControl) Destroy() {
	mc.renderer.Destroy()
	mc.sdl_window.Destroy()
}

func (mc *MissionControl) Start() {
	var event sdl.Event
	var running bool = true

	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\ttype:%d\tid:%d\tx:%d\ty:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				// Quit on escape
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
			}
		}

		// Update each widget
		for _, d := range mc.drawables {
			d.Update()
		}

		// This will update the entire screen to this color
		mc.renderer.SetDrawColor(
			mc.background_color[0],
			mc.background_color[1],
			mc.background_color[2],
			mc.background_color[3])
		mc.renderer.Clear()

		// Update the ui
		for _, d := range mc.drawables {
			d.Draw(mc.renderer)
		}

		mc.renderer.Present()

		// We need to sleep the remainder of the frame here
	}
}

func (mc *MissionControl) Add(d drawableInterface) drawableInterface {
	mc.drawables = append(mc.drawables, d)
	return d
}

func (mc *MissionControl) SetBackgroundColor(r uint8, g uint8, b uint8, a uint8) {
	mc.background_color[0] = r
	mc.background_color[1] = g
	mc.background_color[2] = b
	mc.background_color[3] = a
}

func (mc *MissionControl) SetDefaultBorderColor(r uint8, g uint8, b uint8, a uint8) {
	mc.default_border_color[0] = r
	mc.default_border_color[1] = g
	mc.default_border_color[2] = b
	mc.default_border_color[3] = a
}
