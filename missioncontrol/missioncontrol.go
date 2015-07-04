package missioncontrol

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type MissionControl struct {
	Display  int
	Windows  []*window
	renderer *sdl.Renderer
}

func Init() *MissionControl {
	mc := new(MissionControl)

	var err error

	// Create a new window with a default size
	window, err := sdl.CreateWindow("Mission Control", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}
	defer window.Destroy()

	// Set window to fullscreen
	var disp_rect sdl.Rect
	sdl.GetDisplayBounds(mc.Display, &disp_rect)
	window.SetPosition(int(disp_rect.X), int(disp_rect.Y))
	window.SetSize(int(disp_rect.W), int(disp_rect.H))
	window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	mc.renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	defer mc.renderer.Destroy()

	mc.renderer.Clear()

	return mc
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

	}

	mc.renderer.Clear()

	// Update the ui

	mc.renderer.Present()

	// We need to sleep the remainder of the frame here
}
